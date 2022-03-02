package main

import (
	"bufio"
	"io"
)

func (s *Source) StdioRun(w io.Writer, r io.Reader) error { // флаг -c в оригинальном grep не работает с STDIO
	scanner := bufio.NewScanner(r)
	var index int
	var line string
	furtherMap := make(map[int]int)
	for scanner.Scan() {
		line = scanner.Text()
		s.AddLine(line)
		if s.Config.FlagA > 0 || s.Config.FlagC > 0 { // обработка флагов -A и -C для STDIO
			if v, ok := furtherMap[index]; ok {
				for i := 0; i < v; i++ {
					if err := s.PrintLine(w, index); err != nil {
						return err
					}
				}
			}
		}
		if s.CheckLine(line) {
			if s.Config.FlagB > 0 || s.Config.FlagC > 0 { // обработка флагов -B и -C для STDIO
				if s.Config.FlagB > 0 { // флаг -B имеет больший приоритет, чем -С в оригинальном grep
					if err := s.printPrevious(w, s.Config.FlagB, index); err != nil {
						return err
					}
				} else {
					if err := s.printPrevious(w, s.Config.FlagC, index); err != nil {
						return err
					}
				}
			}
			if err := s.PrintLine(w, index); err != nil {
				return err
			}
			if s.Config.FlagA > 0 || s.Config.FlagC > 0 { // обработка флагов -A и -C для STDIO
				if s.Config.FlagA > 0 { // флаг -A имеет больший приоритет, чем -С в оригинальном grep
					setFurtherMap(furtherMap, s.Config.FlagA, index)
				} else {
					setFurtherMap(furtherMap, s.Config.FlagC, index)
				}
			}
		}
		index++
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}
