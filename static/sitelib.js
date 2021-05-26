$.extend({
    ShowResponseErr: function (result) {
        msg = ""
        if (result) {
            try {
                if (result.message) {
                    return result.message;
                }
                if (result.response.data) {
                    if (result.response.data.message) {
                        msg = result.response.data.message;
                        return msg
                    } else if (result.response.data.error) {
                        msg = result.response.data.error;
                        return msg
                    } else {
                        msg = result.response.data;
                        return msg
                    }
                }
                if (result.message) {
                    msg = result.message;
                    return msg
                }
                var resulObj = JSON.parse(result.responseText)
                if (!resulObj) {
                    msg = result.responseText;
                } else {
                    msg = resulObj.message;
                }
            } catch (e) {
                if (result.statusText) {
                    msg = result.statusText;
                } else {
                    msg = result.responseText;
                }
            }
        }
        if (msg == undefined) {
            msg = ""
        }

        return msg
    },
    getUrlParam: function (name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
        var r = window.location.search.substr(1).match(reg);
        if (r != null)
            return unescape(r[2]);
        return null;
    }
});
