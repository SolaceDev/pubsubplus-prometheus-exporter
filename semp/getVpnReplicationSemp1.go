package semp

import (
	"encoding/xml"
	"errors"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

// Replication Config and status
func (e *Semp) GetVpnReplicationSemp1(ch chan<- prometheus.Metric, vpnFilter string) (ok float64, err error) {
	type Data struct {
		RPC struct {
			Show struct {
				MessageVpn struct {
					Replication struct {
						MessageVpns struct {
							MessageVpn []struct {
								VpnName                    string `xml:"vpn-name"`
								AdminState                 string `xml:"admin-state"`
								ConfigState                string `xml:"config-state"`
								TransactionReplicationMode string `xml:"transaction-replication-mode"`
							} `xml:"message-vpn"`
						} `xml:"message-vpns"`
					} `xml:"replication"`
				} `xml:"message-vpn"`
			} `xml:"show"`
		} `xml:"rpc"`
		ExecuteResult struct {
			Result string `xml:"code,attr"`
		} `xml:"execute-result"`
	}

	command := "<rpc><show><message-vpn><vpn-name>" + vpnFilter + "</vpn-name><replication/></message-vpn></show></rpc>"
	body, err := e.postHTTP(e.brokerURI+"/SEMP", "application/xml", command)
	if err != nil {
		_ = level.Error(e.logger).Log("msg", "Can't scrape VpnSemp1", "err", err, "broker", e.brokerURI)
		return 0, err
	}
	defer body.Close()
	decoder := xml.NewDecoder(body)
	var target Data
	err = decoder.Decode(&target)
	if err != nil {
		_ = level.Error(e.logger).Log("msg", "Can't decode Xml VpnSemp1", "err", err, "broker", e.brokerURI)
		return 0, err
	}
	if target.ExecuteResult.Result != "ok" {
		_ = level.Error(e.logger).Log("msg", "unexpected result", "command", command, "result", target.ExecuteResult.Result, "broker", e.brokerURI)
		return 0, errors.New("unexpected result: see log")
	}

	for _, vpn := range target.RPC.Show.MessageVpn.Replication.MessageVpns.MessageVpn {
		ch <- prometheus.MustNewConstMetric(MetricDesc["VpnReplication"]["vpn_replication_admin_state"], prometheus.GaugeValue, encodeMetricMulti(vpn.AdminState, []string{"shutdown", "enabled", "n/a"}), vpn.VpnName)
		ch <- prometheus.MustNewConstMetric(MetricDesc["VpnReplication"]["vpn_replication_config_state"], prometheus.GaugeValue, encodeMetricMulti(vpn.ConfigState, []string{"standby", "active", "n/a"}), vpn.VpnName)
		ch <- prometheus.MustNewConstMetric(MetricDesc["VpnReplication"]["vpn_replication_transaction_replication_mode"], prometheus.GaugeValue, encodeMetricMulti(vpn.TransactionReplicationMode, []string{"async", "sync", "n/a"}), vpn.VpnName)
	}

	return 1, nil
}
