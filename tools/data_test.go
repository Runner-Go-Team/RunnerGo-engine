package tools

import (
	"fmt"
	"testing"
)

func TestBase64DeEncode(t *testing.T) {
	str := "data:application/pdf;base64,JVBERi0xLjMKJbrfrOAKMyAwIG9iago8PC9UeXBlIC9QYWdlCi9QYXJlbnQgMSAwIFIKL1Jlc291cmNlcyAyIDAgUgovTWVkaWFCb3ggWzAgMCA1OTUuMjc5OTk5OTk5OTk5OTcyNyA4NDEuODg5OTk5OTk5OTk5OTg2NF0KL0NvbnRlbnRzIDQgMCBSCj4+CmVuZG9iago0IDAgb2JqCjw8Ci9MZW5ndGggOTUKPj4Kc3RyZWFtCjAuMjAwMDI1IHcKMCBHCnEKNTk1LjI3OTk5OTk5OTk5OTk3MjcgMCAwIDguMTEzNDI0NjU3NTM0MjQ1OSAwLiA4MzMuNzc2NTc1MzQyNDY1Nzc2IGNtCi9JMCBEbwpRCmVuZHN0cmVhbQplbmRvYmoKMSAwIG9iago8PC9UeXBlIC9QYWdlcwovS2lkcyBbMyAwIFIgXQovQ291bnQgMQo+PgplbmRvYmoKNSAwIG9iago8PAovVHlwZSAvRm9udAovQmFzZUZvbnQgL0hlbHZldGljYQovU3VidHlwZSAvVHlwZTEKL0VuY29kaW5nIC9XaW5BbnNpRW5jb2RpbmcKL0ZpcnN0Q2hhciAzMgovTGFzdENoYXIgMjU1Cj4+CmVuZG9iago2IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9CYXNlRm9udCAvSGVsdmV0aWNhLUJvbGQKL1N1YnR5cGUgL1R5cGUxCi9FbmNvZGluZyAvV2luQW5zaUVuY29kaW5nCi9GaXJzdENoYXIgMzIKL0xhc3RDaGFyIDI1NQo+PgplbmRvYmoKNyAwIG9iago8PAovVHlwZSAvRm9udAovQmFzZUZvbnQgL0hlbHZldGljYS1PYmxpcXVlCi9TdWJ0eXBlIC9UeXBlMQovRW5jb2RpbmcgL1dpbkFuc2lFbmNvZGluZwovRmlyc3RDaGFyIDMyCi9MYXN0Q2hhciAyNTUKPj4KZW5kb2JqCjggMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL0Jhc2VGb250IC9IZWx2ZXRpY2EtQm9sZE9ibGlxdWUKL1N1YnR5cGUgL1R5cGUxCi9FbmNvZGluZyAvV2luQW5zaUVuY29kaW5nCi9GaXJzdENoYXIgMzIKL0xhc3RDaGFyIDI1NQo+PgplbmRvYmoKOSAwIG9iago8PAovVHlwZSAvRm9udAovQmFzZUZvbnQgL0NvdXJpZXIKL1N1YnR5cGUgL1R5cGUxCi9FbmNvZGluZyAvV2luQW5zaUVuY29kaW5nCi9GaXJzdENoYXIgMzIKL0xhc3RDaGFyIDI1NQo+PgplbmRvYmoKMTAgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL0Jhc2VGb250IC9Db3VyaWVyLUJvbGQKL1N1YnR5cGUgL1R5cGUxCi9FbmNvZGluZyAvV2luQW5zaUVuY29kaW5nCi9GaXJzdENoYXIgMzIKL0xhc3RDaGFyIDI1NQo+PgplbmRvYmoKMTEgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL0Jhc2VGb250IC9Db3VyaWVyLU9ibGlxdWUKL1N1YnR5cGUgL1R5cGUxCi9FbmNvZGluZyAvV2luQW5zaUVuY29kaW5nCi9GaXJzdENoYXIgMzIKL0xhc3RDaGFyIDI1NQo+PgplbmRvYmoKMTIgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL0Jhc2VGb250IC9Db3VyaWVyLUJvbGRPYmxpcXVlCi9TdWJ0eXBlIC9UeXBlMQovRW5jb2RpbmcgL1dpbkFuc2lFbmNvZGluZwovRmlyc3RDaGFyIDMyCi9MYXN0Q2hhciAyNTUKPj4KZW5kb2JqCjEzIDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9CYXNlRm9udCAvVGltZXMtUm9tYW4KL1N1YnR5cGUgL1R5cGUxCi9FbmNvZGluZyAvV2luQW5zaUVuY29kaW5nCi9GaXJzdENoYXIgMzIKL0xhc3RDaGFyIDI1NQo+PgplbmRvYmoKMTQgMCBvYmoKPDwKL1R5cGUgL0ZvbnQKL0Jhc2VGb250IC9UaW1lcy1Cb2xkCi9TdWJ0eXBlIC9UeXBlMQovRW5jb2RpbmcgL1dpbkFuc2lFbmNvZGluZwovRmlyc3RDaGFyIDMyCi9MYXN0Q2hhciAyNTUKPj4KZW5kb2JqCjE1IDAgb2JqCjw8Ci9UeXBlIC9Gb250Ci9CYXNlRm9udCAvVGltZXMtSXRhbGljCi9TdWJ0eXBlIC9UeXBlMQovRW5jb2RpbmcgL1dpbkFuc2lFbmNvZGluZwovRmlyc3RDaGFyIDMyCi9MYXN0Q2hhciAyNTUKPj4KZW5kb2JqCjE2IDAgb2JqCjw8Ci9UeXBl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Rml0SCBudWxsXQovUGFnZUxheW91dCAvT25lQ29sdW1uCj4+CmVuZG9iagp4cmVmCjAgMjIKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDAwMjk3IDAwMDAwIG4gCjAwMDAwMjExOTUgMDAwMDAgbiAKMDAwMDAwMDAxNSAwMDAwMCBuIAowMDAwMDAwMTUyIDAwMDAwIG4gCjAwMDAwMDAzNTQgMDAwMDAgbiAKMDAwMDAwMDQ3OSAwMDAwMCBuIAowMDAwMDAwNjA5IDAwMDAwIG4gCjAwMDAwMDA3NDIgMDAwMDAgbiAKMDAwMDAwMDg3OSAwMDAwMCBuIAowMDAwMDAxMDAyIDAwMDAwIG4gCjAwMDAwMDExMzEgMDAwMDAgbiAKMDAwMDAwMTI2MyAwMDAwMCBuIAowMDAwMDAxMzk5IDAwMDAwIG4gCjAwMDAwMDE1MjcgMDAwMDAgbiAKMDAwMDAwMTY1NCAwMDAwMCBuIAowMDAwMDAxNzgzIDAwMDAwIG4gCjAwMDAwMDE5MTYgMDAwMDAgbiAKMDAwMDAwMjAxOCAwMDAwMCBuIAowMDAwMDAyMTE0IDAwMDAwIG4gCjAwMDAwMjE0NTQgMDAwMDAgbiAKMDAwMDAyMTU0MCAwMDAwMCBuIAp0cmFpbGVyCjw8Ci9TaXplIDIyCi9Sb290IDIxIDAgUgovSW5mbyAyMCAwIFIKL0lEIFsgPDRGNEQ1NEQxQUUzQUVDOTY0RjU4Rjc1NkY2QzBGOTI2PiA8NEY0RDU0RDFBRTNBRUM5NjRGNThGNzU2RjZDMEY5MjY+IF0KPj4Kc3RhcnR4cmVmCjIxNjQ0CiUlRU9G"
	_, str = Base64DeEncode(str, "file")
	//fmt.Println(string(code))
	fmt.Println(str)

	//str = Base64Encode("Hello world, 你好, 世界")
	//fmt.Println(str)
	//
	//code = Base64DeEncode(str, "")
	//fmt.Println(string(code))
}
