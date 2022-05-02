package loadgen

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"../log"
	"./lib"
)

var logger = log.Dlogger()

type myGenerator struct {
	caller      lib.Caller
	timeoutNS   time.Duration
	lps         uint32
	durationNS  time.Duration
	concurrency uint32
	tickets     lib.GoTickets //票池
	ctx         context.Context
	cancelFunc  context.CancelFunc
	callCount   int64
	status      uint32
	resultCh    chan *lib.CallResult
}

func NewGenerator(pset ParamSet) (lib.Generator, error) {
	logger.Infoln("New a load generator...")
	if err := pset.Check(); err != nil {
		return nil, err
	}
	gen := &myGenerator{
		caller:     pset.Caller,
		timeoutNS:  pset.TimeoutNS,
		lps:        pset.LPS,
		durationNS: pset.DurationNS,
		status:     lib.STATUS_ORIGINAL,
		resultCh:   pset.ResultCh,
	}
	if err := gen.init(); err != nil {
		return nil, err
	}
	return gen, nil
}

func (gen *myGenerator) init() error {
	var buf bytes.Buffer
	buf.WriteString("Initializing the load generator...")
	var total64 = int64(gen.timeoutNS)/int64(1e9/gen.lps) + 1
	if total64 > math.MaxInt32 {
		total64 = math.MaxInt32
	}
	gen.concurrency = uint32(total64)
	tickets, err := lib.NewGoTickets(gen.concurrency)
	if err != err {
		return err
	}
	gen.tickets = tickets

	buf.WriteString((fmt.Sprintf("Done. (concurrency=%d", gen.concurrency)))
	logger.Infoln(buf.String())
	return nil
}

func (gen *myGenerator) callOne(rawReq *lib.RawReq) *lib.RawResp {
	atomic.AddInt64(&gen.callCount, 1)
	if rawReq == nil {
		return &lib.RawResp{ID: -1, Err: errors.New("Invalid raw request.")}
	}
	start := time.Now().UnixNano()
	resp, err := gen.caller.Call(rawReq.Req, gen.timeoutNS)
	end := time.Now().UnixNano()
	elapsedTime := time.Duration(end - start)
	var rawResp lib.RawResp
	if err != nil{
		errMsg:=fmt.Sprintf("Sync Call Error: %s.", err)
		rawResp = lib.RawResp{
			ID: rawReq.ID,
			Err: errors.New(errMsg),
			Elapse: elapsedTime,
		}
	}else{
		rawResp = lib.RawResp{
			ID: rawReq.ID,
			Resp: resp,
			Elapse: elapsedTime,
		}
	}
	return &rawResp
}

func (gen *myGenerator) asyncCall(){
	gen.tickets.Take()
	go func(){
		defer func(){
			if p:= recover();p!=nil{
				err, ok :=interface{}(p).(error)
				var errMsg string
				if ok{
					errMsg = fmt.Sprintf("Async call panic!(error: %s)",err)
				}else{
					errMsg = fmt.Sprintf("Async call panic! (clue:%#v)",p)
				}
				logger.Errorln(errMsg)
				result := &lib.CallResult{
					ID: -1,
					Code: lib.RET_CODE_FATAL_CALL,
					Msg: errMsg,
				}
				gen.sendResult(result)
			}
		gen.tickets.Return()
	}()
	rawReq := gen.caller.BuildReq()

	var callStatus uint32
	timer
	}
}
