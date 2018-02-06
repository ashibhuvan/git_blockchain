package main


import {

	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
	"time"


	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	}

type Block struct{
	Index	int s
	Timestamp	string
	BPM	int
	Hash	string
	PrevHash	string

	}

var Blockchain []Block



func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodetoString(hashed)

	}

func generateBlock(oldBlock Block, BPM int) (Block, error){

	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil

	}

func isBlockValid(newBlock, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index{
			return false
		}
	if oldBloc.Hash != newBlock.PrevHash{

		return false
	}

	if calculateHash(newBlock) != newBlock.Hash{
		return false
	}


	}

func replaceChain(newBlocks []Block){

	if len(newBlocks) > len(Blockchain){
			Blockchain = newBlocks
		}

	}


func run() error{

	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", os.Getenv("ADDR"))
	s := &http.Server{

		Addr:		":" + httpAddr,
		Handler:		mux,
		ReadTimeout:		10 * time.Second,
		WriteTimeout:		10 * time.Second,
		MaxHeaderBytes:		1 << 20,

	}


	return nil
	}
