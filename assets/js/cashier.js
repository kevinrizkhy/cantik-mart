$('.top.menu .item').tab();
$(document).ready(function() {
    $('#itemList').DataTable({
        scrollY: '50vh',
        scrollCollapse: true,
        paging: false,
        pageLength: 10,
        paging: false,
        searching: false,
        columns: [{
                "width": "10%"
            },
            {
                "width": "25%"
            },
            {
                "width": "15%"
            },
            {
                "width": "5%"
            },
            {
                "width": "15%"
            },
            {
                "width": "15%"
            },
            {
                "width": "15%"
            },
        ]
    });
    $('#itemSearchList').DataTable({
        scrollY: '50vh',
        scrollCollapse: true,
        pageLength: 10,
        searching: false,
        order: [
            [1, "asc"]
        ],
        lengthChange: false
    });
    $('#memberSearchList').DataTable({
        scrollY: '50vh',
        scrollCollapse: true,
        pageLength: 10,
        searching: false,
        lengthChange: false
    });
    document.getElementById("itemId").focus();
    $('.modal').modal();
    $('select').formSelect();
    startConnection();
});


var total = 0;

function PaymentMethod() {
    var method = $('#payment-method-select').val();
    if (method != 1) {
        $("#card").css("display", "block");
        $("#nominal-div").css("display", "none");
        $("#change-div").css("display", "none");
    } else {
        $("#nominal-div").css("display", "block");
        $("#change-div").css("display", "block");
        $("#card").css("display", "none");
    }
}

function exchange() {
    var nominal = parseFloat($('#nominal').val());
    if (total <= nominal) {
        var bilangan = nominal - total;

        var reverse = bilangan.toString().split('').reverse().join(''),
            ribuan = reverse.match(/\d{1,3}/g);
        ribuan = ribuan.join('.').split('').reverse().join('');

        $("#change-h5").html(ribuan);
    }
}

var itemId = document.getElementById('itemId');
var itemSearchId = document.getElementById('itemSearchId');
var memberSearchId = document.getElementById('itemSearchId');
var qtyInput = document.getElementById('item_qty');
var timeout = null;

qtyInput.onkeyup = function(e) {
    clearTimeout(timeout);
    timeout = setTimeout(function() {
        document.getElementById("itemId").focus();
    }, 600);
};

memberSearchId.onkeyup = function(e) {
    clearTimeout(timeout);
    timeout = setTimeout(function() {
        var member = document.getElementById("memberSearchId").value;
        var table = $('#itemSearchList').DataTable();
        $.ajax({
            method: "GET",
            cache: false,
            url: "/api/member-list?search=" + member,
            success: function(result_arr) {
                table.clear().draw();
                for (var i = 0; i < result_arr.length; i++) {
                    var dataSet = [result[0], result[1], result[2], qty, result[3], 100, total];
                    table.row.add(dataSet).draw();
                }
            },
            error: function(error) {
                table.clear().draw();
            }
        });
    }, 300);
};

itemSearchId.onkeyup = function(e) {
    clearTimeout(timeout);
    timeout = setTimeout(function() {
        var item_id = document.getElementById("itemSearchId").value;
        var table = $('#itemSearchList').DataTable();
        $.ajax({
            method: "GET",
            cache: false,
            url: "/api/item-list?search=" + item_id,
            success: function(result_arr) {
                table.clear().draw();
                for (var i = 0; i < result_arr.length; i++) {
                    var result = result_arr[i];
                    var dataSet = [result[0], result[1], result[2], 0, result[3]];
                    table.row.add(dataSet).draw();
                }
            },
            error: function(error) {
                table.clear().draw();
            }
        });
    }, 300);
};

itemId.onkeyup = function(e) {
    clearTimeout(timeout);
    timeout = setTimeout(function() {
        var item_id = document.getElementById("itemId").value;
        $.ajax({
            method: "GET",
            cache: false,
            url: "/api/item?id=" + item_id,
            success: function(result_arr) {
                var table = $('#itemList').DataTable();
                var result = result_arr[0];
                var qty = qtyInput.value;
                var total = parseFloat(qty) * (parseFloat(result[3]) - parseFloat("100"));
                var dataSet = [result[0], result[1], result[2], qty, result[3], 100, total];

                table.row.add(dataSet).draw();
                GrandTotal();

                document.getElementById("itemId").value = "";
                document.getElementById("item_qty").value = 1;
                document.getElementById("itemId").focus();
            },
            error: function(error) {
                document.getElementById("itemId").value = "";
                document.getElementById("item_qty").value = 1;
                document.getElementById("itemId").focus();
            }
        });
    }, 300);
};

function GrandTotal() {
    total = 0;
    var table = $('#itemList').DataTable();
    var data = table.rows().data();
    for (var i = 0; i < data.length; i++) {
        total += parseFloat(data[i][6]);
    }
    var bilangan = total;

    var reverse = bilangan.toString().split('').reverse().join(''),
        ribuan = reverse.match(/\d{1,3}/g);
    ribuan = ribuan.join('.').split('').reverse().join('');

    $("#total-h5").html(ribuan);
    document.getElementById("label-total").innerHTML = ribuan;
}

function CheckOut() {
    alert(123)
}

function PrintBill() {
    printHTML();
}

/// Authentication setup ///
qz.security.setCertificatePromise(function(resolve, reject) {
    //Preferred method - from server
    //        $.ajax("assets/signing/digital-certificate.txt").then(resolve, reject);

    //Alternate method 1 - anonymous
    //        resolve();

    //Alternate method 2 - direct
    resolve("-----BEGIN CERTIFICATE-----\n" +
        "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\n" +
        "-----END CERTIFICATE-----\n" +
        "--START INTERMEDIATE CERT--\n" +
        "-----BEGIN CERTIFICATE-----\n" +
        "MIIFEjCCA/qgAwIBAgICEAAwDQYJKoZIhvcNAQELBQAwgawxCzAJBgNVBAYTAlVT\n" +
        "MQswCQYDVQQIDAJOWTESMBAGA1UEBwwJQ2FuYXN0b3RhMRswGQYDVQQKDBJRWiBJ\n" +
        "bmR1c3RyaWVzLCBMTEMxGzAZBgNVBAsMElFaIEluZHVzdHJpZXMsIExMQzEZMBcG\n" +
        "A1UEAwwQcXppbmR1c3RyaWVzLmNvbTEnMCUGCSqGSIb3DQEJARYYc3VwcG9ydEBx\n" +
        "emluZHVzdHJpZXMuY29tMB4XDTE1MDMwMjAwNTAxOFoXDTM1MDMwMjAwNTAxOFow\n" +
        "gZgxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJOWTEbMBkGA1UECgwSUVogSW5kdXN0\n" +
        "cmllcywgTExDMRswGQYDVQQLDBJRWiBJbmR1c3RyaWVzLCBMTEMxGTAXBgNVBAMM\n" +
        "EHF6aW5kdXN0cmllcy5jb20xJzAlBgkqhkiG9w0BCQEWGHN1cHBvcnRAcXppbmR1\n" +
        "c3RyaWVzLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBANTDgNLU\n" +
        "iohl/rQoZ2bTMHVEk1mA020LYhgfWjO0+GsLlbg5SvWVFWkv4ZgffuVRXLHrwz1H\n" +
        "YpMyo+Zh8ksJF9ssJWCwQGO5ciM6dmoryyB0VZHGY1blewdMuxieXP7Kr6XD3GRM\n" +
        "GAhEwTxjUzI3ksuRunX4IcnRXKYkg5pjs4nLEhXtIZWDLiXPUsyUAEq1U1qdL1AH\n" +
        "EtdK/L3zLATnhPB6ZiM+HzNG4aAPynSA38fpeeZ4R0tINMpFThwNgGUsxYKsP9kh\n" +
        "0gxGl8YHL6ZzC7BC8FXIB/0Wteng0+XLAVto56Pyxt7BdxtNVuVNNXgkCi9tMqVX\n" +
        "xOk3oIvODDt0UoQUZ/umUuoMuOLekYUpZVk4utCqXXlB4mVfS5/zWB6nVxFX8Io1\n" +
        "9FOiDLTwZVtBmzmeikzb6o1QLp9F2TAvlf8+DIGDOo0DpPQUtOUyLPCh5hBaDGFE\n" +
        "ZhE56qPCBiQIc4T2klWX/80C5NZnd/tJNxjyUyk7bjdDzhzT10CGRAsqxAnsjvMD\n" +
        "2KcMf3oXN4PNgyfpbfq2ipxJ1u777Gpbzyf0xoKwH9FYigmqfRH2N2pEdiYawKrX\n" +
        "6pyXzGM4cvQ5X1Yxf2x/+xdTLdVaLnZgwrdqwFYmDejGAldXlYDl3jbBHVM1v+uY\n" +
        "5ItGTjk+3vLrxmvGy5XFVG+8fF/xaVfo5TW5AgMBAAGjUDBOMB0GA1UdDgQWBBSQ\n" +
        "plC3hNS56l/yBYQTeEXoqXVUXDAfBgNVHSMEGDAWgBQDRcZNwPqOqQvagw9BpW0S\n" +
        "BkOpXjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAJIO8SiNr9jpLQ\n" +
        "eUsFUmbueoxyI5L+P5eV92ceVOJ2tAlBA13vzF1NWlpSlrMmQcVUE/K4D01qtr0k\n" +
        "gDs6LUHvj2XXLpyEogitbBgipkQpwCTJVfC9bWYBwEotC7Y8mVjjEV7uXAT71GKT\n" +
        "x8XlB9maf+BTZGgyoulA5pTYJ++7s/xX9gzSWCa+eXGcjguBtYYXaAjjAqFGRAvu\n" +
        "pz1yrDWcA6H94HeErJKUXBakS0Jm/V33JDuVXY+aZ8EQi2kV82aZbNdXll/R6iGw\n" +
        "2ur4rDErnHsiphBgZB71C5FD4cdfSONTsYxmPmyUb5T+KLUouxZ9B0Wh28ucc1Lp\n" +
        "rbO7BnjW\n" +
        "-----END CERTIFICATE-----\n");
});

qz.security.setSignaturePromise(function(toSign) {
    return function(resolve, reject) {
        //Preferred method - from server
        //            $.ajax("/secure/url/for/sign-message?request=" + toSign).then(resolve, reject);

        //Alternate method - unsigned
        resolve();
    };
});


/// Connection ///
function launchQZ() {
    if (!qz.websocket.isActive()) {
        window.location.assign("qz:launch");
        //Retry 5 times, pausing 1 second between each attempt
        startConnection({
            retries: 5,
            delay: 1
        });
    }
}

function startConnection(config) {
    if (!qz.websocket.isActive()) {
        updateState('Waiting', 'default');

        qz.websocket.connect(config).then(function() {
            updateState('Active', 'success');
            findVersion();
        }).catch(handleConnectionError);
    } else {
        displayMessage('An active connection with QZ already exists.', 'alert-warning');
    }
}

function updateState(text, css) {
    $("#launch").show();
}

function getPath() {
    var path = window.location.href;
    return path.substring(0, path.lastIndexOf("/"));
}


/// QZ Config ///
var cfg = null;

function getUpdatedConfig() {
    if (cfg == null) {
        cfg = qz.configs.create(null);
    }

    updateConfig();
    return cfg
}

function updateConfig() {
    var pxlSize = null;
    if ($("#pxlSizeActive").prop('checked')) {
        pxlSize = {
            width: $("#pxlSizeWidth").val(),
            height: $("#pxlSizeHeight").val()
        };
    }

    var pxlMargins = $("#pxlMargins").val();
    if ($("#pxlMarginsActive").prop('checked')) {
        pxlMargins = {
            top: $("#pxlMarginsTop").val(),
            right: $("#pxlMarginsRight").val(),
            bottom: $("#pxlMarginsBottom").val(),
            left: $("#pxlMarginsLeft").val()
        };
    }

    var copies = 1;
    var jobName = null;
    if ($("#rawTab").hasClass("active")) {
        copies = $("#rawCopies").val();
        jobName = $("#rawJobName").val();
    } else {
        copies = $("#pxlCopies").val();
        jobName = $("#pxlJobName").val();
    }

    cfg.reconfigure({
        altPrinting: $("#rawAltPrinting").prop('checked'),
        encoding: $("#rawEncoding").val(),
        endOfDoc: $("#rawEndOfDoc").val(),
        perSpool: $("#rawPerSpool").val(),

        colorType: $("#pxlColorType").val(),
        copies: copies,
        density: $("#pxlDensity").val(),
        duplex: $("#pxlDuplex").prop('checked'),
        interpolation: $("#pxlInterpolation").val(),
        jobName: jobName,
        legacy: $("#pxlLegacy").prop('checked'),
        margins: pxlMargins,
        orientation: $("#pxlOrientation").val(),
        paperThickness: $("#pxlPaperThickness").val(),
        printerTray: $("#pxlPrinterTray").val(),
        rasterize: $("#pxlRasterize").prop('checked'),
        rotation: $("#pxlRotation").val(),
        scaleContent: $("#pxlScale").prop('checked'),
        size: pxlSize,
        units: $("input[name='pxlUnits']:checked").val()
    });
}

function setPrinter(printer) {
    var cf = getUpdatedConfig();
    cf.setPrinter(printer);

    if (typeof printer === 'object' && printer.name == undefined) {
        var shown;
        if (printer.file != undefined) {
            shown = "<em>FILE:</em> " + printer.file;
        }
        if (printer.host != undefined) {
            shown = "<em>HOST:</em> " + printer.host + ":" + printer.port;
        }

        $("#configPrinter").html(shown);
    } else {
        if (printer.name != undefined) {
            printer = printer.name;
        }

        if (printer == undefined) {
            printer = 'NONE';
        }
        $("#configPrinter").html(printer);
    }
}