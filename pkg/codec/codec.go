package codec

import (
	"io"

	"github.com/pion/mediadevices/pkg/io/audio"
	"github.com/pion/mediadevices/pkg/io/video"
	"github.com/pion/mediadevices/pkg/prop"
)

type VideoEncoderBuilder func(r video.Reader, p prop.Video) (io.ReadCloser, error)
type AudioEncoderBuilder func(r audio.Reader, p prop.Audio) (io.ReadCloser, error)
