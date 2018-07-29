(function (window, undefind) {
    function imgUpLoad(options) {
        var defaults = {
            inputId: "imageFile",        //输入框ID
            formId: "uploadForm",        //表单ID
            paramName: "请选择图片",        //输入框的name
            requestUrl: "",     //请求的URL
            imgSizeScale: 3,     //图片缩放比例
            imgQuality: 0.3,    //图片质量
            sucFun: undefind, //成功的回调函数
            errFun: undefind, //失败的回调函数
            abortFun: undefind     //上传取消的回调函数
        };
        console.log(defaults);
        var opts = $.extend(true, defaults, options || {}),
            _this = document.getElementById(opts.inputId);
        _this.addEventListener('change', fileChange, false);

        function fileChange() {
            var file = _this.files[0];
            var rFilter = /^(image\/bmp|image\/gif|image\/jpeg|image\/png|image\/tiff)$/i;
            if (!rFilter.test(file.type)) {
                alert("文件格式必须为图片");
                return;
            }
            var reader = new FileReader(), image = new Image(),
                canvas = document.createElement("canvas"), ctx = canvas.getContext("2d");
            startFileLoad(reader, image, canvas, ctx);
        }

        function startFileLoad(reader, image, canvas, ctx) {    //文件加载
            reader.onload = function () {        //文件加载完成
                var url = reader.result;
                image.src = url;
            };
            image.onload = function () {        //图片加载完成
                var w = image.naturalWidth, h = image.naturalHeight;
                canvas.width = w / opts.imgSizeScale;
                canvas.height = h / opts.imgSizeScale;
                ctx.drawImage(image, 0, 0, w, h,
                    , 0, canvas.width, canvas.height);
                fileUpload(canvas);
            };
            reader.readAsDataURL(file);
        }

        function fileUpload(canvas) {        //文件上传
            var blob = getBlob(canvas);
            var fd = new FormData(document.getElementById(opts.formId));
            fd.append(opts.paramName, blob, "upload.jpg");
            var xhr = new XMLHttpRequest();
            xhr.addEventListener('load', function (resUpload) {    //请求成功
                _this.style.display = "";
                if (opts.sucFun && typeof opts.sucFun === "function") opts.sucFun(resUpload.currentTarget.response);
            }, false);
            xhr.addEventListener('error', function () {    //请求失败
                _this.style.display = "";
                if (opts.errFun && typeof opts.errFun === "function") opts.errFun();
            }, false);
            xhr.addEventListener('abort', function () {    //上传终止
                _this.style.display = "";
                if (opts.abortFun && typeof opts.abortFun === "function") opts.abortFun();
            }, false);
            xhr.open('POST', opts.requestUrl);//请求地址
            xhr.send(fd);//发送
        }

        function getBlob(canvas) {        //获取blob对象
            var data = canvas.toDataURL("image/jpeg", opts.imgQuality);
            data = data.split(',')[1];
            data = window.atob(data);
            var ia = new Uint8Array(data.length);
            for (var i = 0; i < data.length; i++) {
                ia[i] = data.charCodeAt(i);
            }
            return new Blob([ia], {
                type: "image/jpeg"
            });
        }
    }

    window.imgUpLoad = imgUpLoad;
})(window);

var iMaxFilesize = 2097152; //2M
window.fileSelected = function () {
    var oFile = document.getElementById('imageFile').files[0];    //读取文件
    var rFilter = /^(image\/bmp|image\/gif|image\/jpeg|image\/png|image\/tiff)$/i;
    if (!rFilter.test(oFile.type)) {
        alert("文件格式必须为图片");
        return;
    }
    if (oFile.size > iMaxFilesize) {
        alert("图片大小不能超过2M");
        return;
    }
    var vFD = new FormData(document.getElementById('uploadForm')),    //建立请求和数据
        oXHR = new XMLHttpRequest();
    oXHR.addEventListener('load', function (resUpload) {
        //成功
    }, false);
    oXHR.addEventListener('error', function () {
        //失败
    }, false);
    oXHR.addEventListener('abort', function () {
        //上传中断
    }, false);
    oXHR.open('POST', actionUrl);
    oXHR.send(vFD);
};