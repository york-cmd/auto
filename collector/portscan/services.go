package portscan

import (
	"auto/collector/cdn"
	"auto/models"
	"auto/utils"
	"fmt"
	"strings"
)

var (
	urls           []string
	ipWebPorts     = make(map[string][]string)
	serviceIPPorts = make(map[string][]string)
)
var (
	ServicePath = "result/services"
)
var webDefaultProtocolPort = []models.ProtocolPort{
	{
		Protocol: "http",
		Port:     "",
	},
	{
		Protocol: "https",
		Port:     "",
	},
	{
		Protocol: "http",
		Port:     "4848",
	},
	{
		Protocol: "http",
		Port:     "7070",
	},
	{
		Protocol: "http",
		Port:     "8080",
	},
	{
		Protocol: "http",
		Port:     "8089",
	},
	{
		Protocol: "https",
		Port:     "8181",
	},
	{
		Protocol: "https",
		Port:     "8443",
	},
	{
		Protocol: "http",
		Port:     "8888",
	},
	{
		Protocol: "http",
		Port:     "9080",
	},
	{
		Protocol: "https",
		Port:     "9443",
	},
	{
		Protocol: "http",
		Port:     "9060",
	},
	{
		Protocol: "http",
		Port:     "9043",
	},
}

func getServices() {
	getWebService()
	getOtherServices()
}
func getWebService() {
	getNoCDNWebService()
	getYesCDNWebServices()
	getFirewallWebServices()
	utils.SliceWriter(ServicePath+"/urls.txt", urls)
}
func getNoCDNWebService() {
	for ip, portServices := range IPPortServicesMap {
		for _, portService := range portServices {
			if strings.Contains(portService.Service, "http") {
				ipWebPorts[ip] = append(ipWebPorts[ip], portService.Port)
			}
		}
	}
	for ip, ports := range ipWebPorts {
		for _, domain := range cdn.IPDomainsMap[ip] {
			for _, port := range ports {
				urls = append(urls, getUrl(domain, port))
			}
		}
	}
}
func getYesCDNWebServices() {
	for _, domain := range cdn.YesCDNDomains {
		urls = append(urls, getWebDefaultUrls(domain)...)
	}
}
func getFirewallWebServices() {
	for _, ip := range YesFirewallIPS {
		for _, domain := range cdn.IPDomainsMap[ip] {
			urls = append(urls, getWebDefaultUrls(domain)...)
		}
	}
}
func getOtherServices() {
	for ip, portServices := range IPPortServicesMap {
		for _, data := range portServices {
			serviceIPPorts[data.Service] = append(serviceIPPorts[data.Service], ip+":"+data.Port)
		}
	}
	for service, data := range serviceIPPorts {
		utils.SliceWriter(ServicePath+"/"+service+".txt", data)
	}
}
func getWebDefaultUrls(domain string) []string {
	var urls []string
	for _, data := range webDefaultProtocolPort {
		if data.Port != "" {
			urls = append(urls, fmt.Sprintf("%s://%s:%s", data.Protocol, domain, data.Port))
		} else {
			urls = append(urls, fmt.Sprintf("%s://%s", data.Protocol, domain))
		}
	}
	return urls
}
func getUrl(domain, port string) string {
	if port == "443" {
		return "https://" + domain
	}
	if port == "80" {
		return "http://" + domain
	}
	if port == "8443" {
		return "https://" + domain + ":8443"
	}
	return "http://" + domain + ":" + port
}
