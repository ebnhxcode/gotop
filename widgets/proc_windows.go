package widgets

import (
	"fmt"
	"log"

	psProc "github.com/shirou/gopsutil/process"
)

func getProcs() ([]Proc, error) {
	psProcs, err := psProc.Processes()
	if err != nil {
		return nil, fmt.Errorf(tr.Value("widget.proc.err.gopsutil", err.Error()))
	}

	procs := make([]Proc, len(psProcs))
	for i, psProc := range psProcs {
		pid := psProc.Pid
		command, err := psProc.Name()
		if err != nil {
			log.Println(tr.Value("widget.proc.err.getcmd", err, psProc, i, pid))
		}
		cpu, err := psProc.CPUPercent()
		if err != nil {
			log.Println(tr.Value("widget.proc.err.cpupercent", err, psProc, i, pid))
		}
		mem, err := psProc.MemoryPercent()
		if err != nil {
			log.Println(tr.Value("widget.proc.err.mempercent", err, psProc, i, pid))
		}

		procs[i] = Proc{
			Pid:         int(pid),
			CommandName: command,
			CPU:         cpu,
			Mem:         float64(mem),
			// getting command args using gopsutil's Cmdline and CmdlineSlice wasn't
			// working the last time I tried it, so we're just reusing 'command'
			FullCommand: command,
		}
	}

	return procs, nil
}