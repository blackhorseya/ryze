package kafkax

import (
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// WriterOptions defines the options for writer
type WriterOptions struct {
	Brokers  []string `json:"brokers" yaml:"brokers"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
}

// NewWriterOptions serve caller to get a new WriterOptions instance
func NewWriterOptions(v *viper.Viper, logger *zap.Logger) (*WriterOptions, error) {
	o := new(WriterOptions)
	err := v.UnmarshalKey("kafka", o)
	if err != nil {
		return nil, err
	}

	logger.Info("get kafka writer options success")

	return o, nil
}

// NewWriter serve caller to get a new Writer instance
func NewWriter(o *WriterOptions) (*kafka.Writer, error) {
	mechanism := plain.Mechanism{
		Username: o.Username,
		Password: o.Password,
	}

	w := &kafka.Writer{
		Addr:     kafka.TCP(o.Brokers...),
		Balancer: &kafka.Hash{},
		Transport: &kafka.Transport{
			SASL: mechanism,
		},
	}

	return w, nil
}

// WriterSet declare the provider set for writer
var WriterSet = wire.NewSet(NewWriterOptions, NewWriter)
