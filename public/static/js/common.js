layui.use(['element', 'layer'], function () {
    var $ = layui.jquery;
    //ajax封装
    /*
     *Ajax 请求权限异常
     *   用户权限错误跳转登陆页
     *   404错误跳转404页面
     */
    $(document).ajaxComplete(function (evt, req, settings) {
        if (req && req.responseJSON) {
            if (req.responseJSON.code == 10000) {
                location.href = '/admin/login';
            }
        }
    });
    /*
     *Ajax 请求错误提示
     *例如：500错误
     *返回错误信息格式
     *{
     *   code: 500,
     *   info: 系统发生异常
     *}
     */
    $(document).ajaxError(function (evt, req, settings) {
        if (req && (req.status === 200 || req.status === 0)) {
            return false;
        }
        var msg = '错误：';
        if (req && req.responseJSON) {
            var json = req.responseJSON;
            msg += json.code || '';
            msg += json.info || '系统异常，请重试';
        } else {
            msg = '系统异常，请重试';
        }
        layer.alert(msg);
    });

    window.PH = {};
    /**
     * ajax封装
     */
    PH.api = function (url, data, callback) {
        // ajax 请求参数
        var ajaxSettings = function (opt) {
            var url = opt.url;
            var href = location.href;
            // 判断是否跨域请求
            var requestType = 'json';
            /*if (url.indexOf(location.host) > -1)
             requestType = 'json';*/
            requestType = opt.dataType || requestType;
            // 是否异步请求
            var async = (opt.async === undefined ? true : opt.async);
            var loadingIndex = layer.load(1, {
                shade: [0.5, '#000'] //0.1透明度的白色背景
            });
            return {
                url: url,
                async: async,
                type: opt.type || 'post',
                dataType: requestType,
                cache: false,
                // xhrFields: {
                //     withCredentials: true
                // },
                // crossDomain: true,
                data: $.extend(opt.data, {sessionId: $.cookie('sessionId')}),
                success: function (data, textStatus, xhr) {
                    if ((requestType === 'json' || requestType === "jsonp") && typeof (data) === "string") {
                        data = JSON.parse(data);
                    }
                    if (data.code != 0) {
                        layer.alert(data.message);
                        return;
                    }
                    opt.success(data);
                },
                error: function (xhr, status, handler) {
                    if (opt.error)
                        opt.error();
                },
                complete: function () {
                    layer.close(loadingIndex);
                }
            };
        };
        $.ajax(ajaxSettings({
            url: url,
            data: data,
            success: function (ret) {
                callback(ret);
            }
        }));
    };

    PH.api2 = function (url, data, success, async) {
        //var url, data, success;
        var aLen = arguments.length;
        if (aLen == 2) {
            url = arguments[0];
            success = data = arguments[1];
        } else if (aLen == 3) {
            url = arguments[0];
            data = arguments[1];
            success = arguments[2];
        } else if (aLen == 4) {

        } else {
            throw "arguments length should be 2 or 3";
        }
        if (aLen < 4) {
            async = true;
        }
        var loadingIndex = layer.load(1, {
            shade: [0.5, '#000'] //0.1透明度的白色背景
        });
        $.ajax({
            type: "post",
            url: url,
            data: data,
            async: async,
            dataType: 'json',
            success: function (ret) {
                if (ret.code != 0) {
                    layer.alert(ret.message);
                    return;
                }
                success && success(ret);
            },
            error: function (e) {
                layer.alert(e);
            },
            complete: function () {
                layer.close(loadingIndex);
            }
        });
    };

    /**
     * 获取url参数
     */
    $.extend({
        getUrlVars: function () {
            var vars = [], hash;
            var url = location.hash;
            var hashes = url.slice(url.indexOf('?') + 1).split('&');
            for (var i = 0; i < hashes.length; i++) {
                hash = hashes[i].split('=');
                vars.push(hash[0]);
                vars[hash[0]] = hash[1];
            }
            return vars;
        },
        getUrlVar: function (name) {
            return $.getUrlVars()[name];
        }
    });

    PH.randomString = function (len) {
        len = len || 32;
        var $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678';
        /****默认去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1****/
        var maxPos = $chars.length;
        var pwd = '';
        for (i = 0; i < len; i++) {
            pwd += $chars.charAt(Math.floor(Math.random() * maxPos));
        }
        return pwd;
    };

    /**
     * 时间格式化
     * @returns {format}
     * @constructor
     */
    Date.prototype.format = function (format) {
        var o = {
            "M+": this.getMonth() + 1, // month
            "d+": this.getDate(), // day
            "h+": this.getHours(), // hour
            "m+": this.getMinutes(), // minute
            "s+": this.getSeconds(), // second
            "q+": Math.floor((this.getMonth() + 3) / 3), // quarter
            "S": this.getMilliseconds()
        }
        if (/(y+)/.test(format))
            format = format.replace(RegExp.$1, (this.getFullYear() + "")
                .substr(4 - RegExp.$1.length));
        for (var k in o)
            if (new RegExp("(" + k + ")").test(format))
                format = format.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length));
        return format;
    };

    PH.table = function (selector, url, params, columns) {
        if (arguments.length == 3) {
            columns = arguments[2];
            params = {};
        }
        return $(selector).DataTable({
            "serverSide": true,
            "processing": true,
            "ordering": false,
            "searching": false,
            "autoWidth": false,
            "scrollX": true,
            "lengthMenu": [100, 50, 10],
            "select":true,
            // "select": {
            //     style: 'os',
            //     // selector: 'td:first-child',
            // },
            "ajax": {
                "url": url,
                "data": function (d) {
                    let newData = {};
                    newData.q = d.search.value;
                    newData.page = d.start / d.length + 1;
                    newData.pageSize = d.length;
                    if (typeof (params) === "function") {
                        return $.extend(newData, params());
                    }
                    return $.extend(newData, params);
                },
            },
            "language": {
                "paginate": {
                    "first": "首页",
                    "last": "尾页",
                    "next": "下一页",
                    "previous": "上一页",
                },
                "search": "搜索:",
                "infoFiltered": "",
                "info": "显示第 _PAGE_ 页,总共: _PAGES_ 页",
                "lengthMenu": "每页显示 _MENU_ 条数据",
                "loadingRecords": "加载中...",
                "processing": "正在处理...",
                "zeroRecords": "数据为空",
                "emptyTable": "数据为空",
            },
            "columns": columns,
        });
    }
});