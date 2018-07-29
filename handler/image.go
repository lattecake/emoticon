package handler

import (
	"net/http"
	"log"
	"github.com/lattecake/emoticon/core"
	"strconv"
	"time"
	"os"
	"image"
	"image/jpeg"
	"bytes"
	"io/ioutil"
)

type ImageHandler struct {
	Path string
}

func (c *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("uri", r.RequestURI, "addr", r.RemoteAddr)
	if r.Method != "POST" {
		http.Redirect(w, r, "/index.html", http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	file, head, err := r.FormFile("imageUpload")
	if err != nil {
		log.Println("r", "FormFile", "err", err.Error())
		w.Write(core.Response(false, err.Error(), nil))
		return
	}

	defer file.Close()

	firefix := strconv.FormatInt(time.Now().Unix(), 10)
	filePath := c.Path + time.Now().Format("/2006/01/02/15/")

	if _, err := PathExistsAndMkdir(filePath); err != nil {
		log.Println("func", "PathExistsAndMkdir", "err", err.Error())
		w.Write(core.Response(false, err.Error(), nil))
		return
	}

	fw, err := os.Create(filePath + core.Sign(head.Filename+firefix) + ".jpg")
	if err != nil {
		log.Println(err.Error())
		w.Write(core.Response(false, "文件创建失败"+err.Error(), nil))
		return
	}
	defer fw.Close()

	d, err := ioutil.ReadAll(fw)
	if err != nil {
		log.Println("ioutil", "readall", "err", err.Error())
		w.Write(core.Response(false, "文件创建失败"+err.Error(), nil))
		return
	}

	m, _, err := image.Decode(bytes.NewBuffer(d))
	if err != nil {
		log.Println("image", "decode", "err", err.Error())
		w.Write(core.Response(false, err.Error(), nil))
		return
	}
	rgbImg := m.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(0, 0, 200, 200)).(*image.YCbCr)

	f, err := os.Create(filePath + core.Sign(head.Filename+firefix) + "-m.jpg")
	if err != nil {
		log.Println("os", "create", "err", err.Error())
		w.Write(core.Response(false, err.Error(), nil))
		return
	}
	defer f.Close()

	if err = jpeg.Encode(f, subImg, nil); err != nil {
		log.Println("jpeg", "encode", "err", err.Error())
		w.Write(core.Response(false, "图片创建失败"+err.Error(), nil))
		return
	}

	core.Response(true, "", map[string]interface{}{
		"image_url": "",
	})

}

func PathExistsAndMkdir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err == nil {
			return true, err
		}
		return false, nil
	}
	return false, err
}
