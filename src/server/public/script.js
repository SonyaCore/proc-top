(function () {
  var g = {}; // graphs
  var util = {}; // utility functions
  util.padt = function (s) {
    return ("00" + s).slice(-2);
  };
  util.seconds = function (n) {
    var d = Math.floor(n / 86400);
    var h = Math.floor(n / 3600) % 24;
    var m = Math.floor(n / 60) % 60;
    var s = Math.floor(n) % 60;
    return [d, util.padt(h), util.padt(m), util.padt(s)].join(":");
  };
  util.bytes = function (n, decimal) {
    var base = decimal ? 10 : 2;
    var exp = decimal ? 3 : 10;
    var units = decimal
      ? ["B", "KB", "MB", "GB", "TB", "PB"]
      : ["B", "KiB", "MiB", "GiB", "TiB", "PiB"];
    if (n < 0) {
      n = -n;
      s = "-";
    } else {
      s = "";
    }
    for (i = 5; i >= 0; i--)
      if (n >= Math.pow(base, i * exp) - 1)
        return s + (n / Math.pow(base, i * exp)).toFixed(2) + " " + units[i];
  };
  util.percent = function (n) {
    return n.toFixed(1) + "%";
  };
  var handler = {}; // data handlers
  handler.fqdn = function (s) {
    $("#fqdn").text(s);
  };
  handler.uptime = function (n) {
    $("#uptime").text(util.seconds(n));
  };
  handler.cpuusage = function (n) {
    $("#cpuusage").text(util.percent(n));
    g.cpuusage.t.append(+new Date(), n);
  };
  handler.ramusage = function (n) {
    $("#ramusage").text(util.bytes(n[0] - n[1]));
    g.ramusage.t.append(+new Date(), n[2]);
  };
  handler.diskio = function (n) {
    $("#diskr").text(util.bytes(n[0], 1));
    $("#diskw").text(util.bytes(n[1], 1));
    if (handler.diskio.lastr != undefined) {
      var rs = n[2] - (handler.diskio.lastr || 0);
      var ws = n[3] - (handler.diskio.lastw || 0);
      $("#diskrs").text(util.bytes(rs, 1) + "/s");
      $("#diskws").text(util.bytes(ws, 1) + "/s");
      g.diskrs.t.append(+new Date(), rs / 1048576);
      g.diskws.t.append(+new Date(), ws / 1048576);
    }
    handler.diskio.lastr = n[2];
    handler.diskio.lastw = n[3];
  };
  handler.diskusage = function (n) {
    $("#disku").text(util.bytes(n[0], 1));
    $("#diskt").text(util.bytes(n[1], 1));
  };
  handler.netio = function (n) {
    $("#nett").text(util.bytes(n[0], 1));
    $("#netr").text(util.bytes(n[1], 1));
    if (handler.netio.lastt != undefined) {
      var ts = n[0] - (handler.netio.lastt || 0);
      var rs = n[1] - (handler.netio.lastr || 0);
      $("#netts").text(util.bytes(ts, 1) + "/s");
      $("#netrs").text(util.bytes(rs, 1) + "/s");
      g.netts.t.append(+new Date(), ts / 1048576);
      g.netrs.t.append(+new Date(), rs / 1048576);
    }
    handler.netio.lastt = n[0];
    handler.netio.lastr = n[1];
  };
  handler.swapusage = function (n) {
    $("#swapusage").text(util.bytes(n[1]));
    g.swapusage.t.append(+new Date(), n[3]);
  };
  handler.kernel = function (s) {
    $("#kernel").text(s);
  };
  handler.loadaverage = function (s) {
    $("#loadaverage").text(s);
  };
  var count = 0,
    errors = 0;
  var latency = 0;
  var wait = 1000;
  var margin = 250;
  function ping() {
    var time = +new Date();
    $.get("raw", function (data) {
      document.title = data.fqdn;
      for (var i in data) handler[i] && handler[i](data[i]);
      var t = (latency = new Date() - time);
      ++count;
      update();
      if (t <= margin) {
        setTimeout(margin - t);
        setTimeout(ping, wait - t);
      } else if (t <= wait - margin) {
        setTimeout(ping, wait - t);
      } else {
        setTimeout(ping, margin);
      }
    });
  }
  function update() {
    $("#latency").text(latency + " ms");
    g.latency.t.append(+new Date(), latency);
    $("#requests").text(count + "/" + errors);
  }
  function error() {
    ++errors;
    $("#uptime").text("OFFLINE");
    update();
    setTimeout(ping, wait);
  }
  function graph(name, percentage) {
    var options = percentage
      ? {
          millisPerPixel: 100,
          grid: {
            fillStyle: "transparent",
            strokeStyle: "rgba(0, 0, 0, 0)",
          },
          minValue: 0,
          maxValue: 100,
          labels: { fillStyle: "rgba(0, 0, 0, 0)" },
        }
      : {
          millisPerPixel: 100,
          grid: {
            fillStyle: "transparent",
            strokeStyle: "rgba(0, 0, 0, 0)",
          },
        };
    var ts_options = {
      strokeStyle: "#00ff00",
    };
    g[name] = {
      c: new SmoothieChart(options),
    };
    g[name].c.streamTo($("#c_" + name)[0], 1000);
    g[name].t = new TimeSeries();
    g[name].c.addTimeSeries(g[name].t, ts_options);
  }
  $.ajaxSetup({
    timeout: 5000,
    error: error,
  });
  var graphlist = {
    latency: 0,
    cpuusage: 1,
    ramusage: 1,
    diskrs: 0,
    diskws: 0,
    netts: 0,
    netrs: 0,
    swapusage: 1,
  };
  for (var i in graphlist) graph(i, graphlist[i]);
  ping();
})();
