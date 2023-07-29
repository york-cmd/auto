package urlvalid

import "log"

func Run() {
	log.Println("start using httpx to verify and deduplicate url ...")
	httpxRun()
	httpxResult()
	log.Println("completed url verification and deduplication .")
}
