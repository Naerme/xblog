package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/ai_service"
	"fmt"
)

//func chat() {
//	url := "https://api.chatanywhere.tech/v1/chat/completions"
//	method := "POST"
//
//	payload := strings.NewReader(`{
//    "model": "gpt-3.5-turbo",
//    "messages": [
//      {
//        "role": "system",
//        "content": "你是一个叫枫枫知道的人工智能助手"
//      },
//      {
//        "role": "user",
//        "content": "你好，你是谁?"
//      }
//    ]
//  }`)
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, payload)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", global.Conifg.Ai.SecretKey))
//	req.Header.Add("Content-Type", "application/json")
//
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//}

//type Choice struct {
//	Index int `json:"index"`
//	Delta struct {
//		Content string `json:"content"`
//	} `json:"delta"`
//	Logprobs     interface{} `json:"logprobs"`
//	FinishReason interface{} `json:"finish_reason"`
//}

type Choices struct {
	Index int `json:"index"`
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason interface{} `json:"finish_reason"`
}

type StreamData struct {
	Id                string    `json:"id"`
	Choices           []Choices `json:"choices"`
	Created           int       `json:"created"`
	Model             string    `json:"model"`
	Object            string    `json:"object"`
	SystemFingerprint string    `json:"system_fingerprint"`
}

//type StreamData struct {
//	Id                string      `json:"id"`
//	Choices           []Choice    `json:"choices"`
//	Created           int         `json:"created"`
//	Model             string      `json:"model"`
//	Object            string      `json:"object"`
//	SystemFingerprint interface{} `json:"system_fingerprint"`
//}

//	func chatStream() {
//		url := "https://api.chatanywhere.tech/v1/chat/completions"
//		method := "POST"
//
//		payload := strings.NewReader(`{
//	   "model": "gpt-3.5-turbo",
//	   "messages": [
//	     {
//	       "role": "system",
//	       "content": "你是一个叫枫枫知道的人工智能助手"
//	     },
//	     {
//	       "role": "user",
//	       "content": "你好，你是谁?"
//	     }
//	   ],
//
// "stream": true
//
//	 }`)
//
//		client := &http.Client{}
//		req, err := http.NewRequest(method, url, payload)
//
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", global.Conifg.Ai.SecretKey))
//		req.Header.Add("Content-Type", "application/json")
//
//		res, err := client.Do(req)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer res.Body.Close()
//
//		scanner := bufio.NewScanner(res.Body)
//		// 按行分割
//		scanner.Split(bufio.ScanLines)
//
//		for scanner.Scan() {
//			text := scanner.Text()
//			if text == "" {
//				continue
//			}
//
//			data := text[6:]
//			if data == "[DONE]" {
//				break
//			}
//			var item StreamData
//			err := json.Unmarshal([]byte(data), &item)
//			if err != nil {
//				fmt.Printf("shibai %s %s", err, data)
//				continue
//			}
//			//fmt.Println("%q\n", text)
//			//fmt.Printf(item.Choices[0].Delta.Content)
//			if len(item.Choices) > 0 && item.Choices[0].Delta.Content != "" {
//				fmt.Print(item.Choices[0].Delta.Content)
//			}
//			if len(item.Choices) == 0 {
//				fmt.Printf("⚠️  choices 为空，原始 data: %s\n", data)
//				continue
//			}
//
//		}
//		fmt.Println()
//
// }
func main() {
	flags.Parse()
	global.Conifg = core.ReadConf()
	core.InitLogrus()

	//chat()

	//chatStream()

	//msg, err := ai_service.Chat("1. Redis的安装与基础配置\n在重启Redis之前，首先需要确保Redis已在Windows系统上正确安装。可以从[Redis的官方网站](\n\n在安装完成后，您可以通过以下命令检查Redis是否成功运行：\n\n登录后复制\nredis-server --version\n1.\n若返回版本号，则表示Redis已成功安装。\n\n2. Redis服务的启动\n在Windows系统中，可以通过命令行启动Redis服务。请打开命令提示符并输入以下命令：\n\n登录后复制\ncd C:\\Program Files\\Redis\nredis-server.exe redis.windows.conf\n1.\n2.\n这将基于 redis.windows.conf 配置文件启动Redis服务。\n\n3. Redis服务的重启\n3.1 使用命令行重启\nRedis本身并没有直接的重启命令，但我们可以通过停止并重新启动服务的方式来实现。可以使用以下步骤：\n\n停止Redis服务\n首先，在命令提示符中输入以下命令停止Redis服务：\n\n登录后复制\nredis-cli shutdown\n1.\n重新启动Redis服务\n然后，再次使用以下命令启动Redis服务：\n\n登录后复制\nredis-server.exe redis.windows.conf\n1.\n这个简单的操作步骤可以实现Redis服务的重启。\n\n3.2 使用PowerShell重启\n如果您习惯使用PowerShell，也可以在PowerShell中执行以下命令：\n\n停止Redis服务：\n登录后复制\nStop-Process -Name \"redis-server\"\n1.\n启动Redis服务：\n登录后复制\nStart-Process \"C:\\Program Files\\Redis\\redis-server.exe\" \"C:\\Program Files\\Redis\\redis.windows.conf\"\n1.\n使用PowerShell进行管理可以使得操作更加灵活且易于自动化。\n\n4. 监控Redis服务状态\n在重启服务之前，监控Redis服务的状态是非常重要的。一种常用的方式是使用命令行工具 redis-cli 来查看Redis的运行情况。\n\n可以运行以下命令查看Redis的状态：\n\n登录后复制\nredis-cli ping\n1.\n如果返回结果是 PONG，说明服务正常。如果服务未响应，则需要检查服务日志。\n\n5. 使用饼状图分析Redis使用情况\n在重启Redis之前，了解Redis的内存及请求情况十分重要。下图以饼状图的形式展示了一些常用的Redis操作占比：\n\n40%\n30%\n15%\n15%\nRedis操作占比\nGET操作\nSET操作\nDEL操作\n其他\n通过分析这些数据，可以更好地调整Redis的配置，以便在重启后能够优化性能。\n\n6. 结尾\n重启Redis在开发和生产环境中都是一个常见的操作。无论是通过命令行还是PowerShell方式，理解这一过程不仅帮助我们维护服务的稳定性，也为故障排查提供了便捷的途径。\n\n在本文中，我们详细介绍了在Windows系统上如何重启Redis服务，包括相关的代码示例和操作步骤。希望这个方案能够帮助您在实际操作中更顺利地进行Redis的管理与维护。\n\n如有需要，更深入的详细分析或复杂场景下的重启方案，欢迎进一步讨论与交流。")
	//msg, err := ai_service.Chat("nihao")
	//fmt.Println(msg, err)

	msgChan, err := ai_service.ChatStream("关于环境搭建的文章")
	if err != nil {
		fmt.Println(err)
	}
	for s := range msgChan {
		fmt.Printf(s)
	}

}
