package solvers

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day5 struct {
	seeds []int
	maps  map[string]*seedMap
}

func (d *Day5) Part1(in []byte) int {
	d.parse(toLines(in))

	lowestLocation := math.MaxInt
	for _, s := range d.seeds {
		l := d.getLocation(s)
		if l < lowestLocation {
			lowestLocation = l
		}
	}
	return int(lowestLocation)
}

func (d *Day5) Part2(in []byte) int {
	d.parse(toLines(in))

	lowestLocation := math.MaxInt
	seeds := d.getSeeds()
	fmt.Println("Testing with", len(seeds), "seeds")
	for i, s := range seeds {
		if i%10000 == 0 {
			fmt.Printf("testing seed %v of %v\n", i, len(seeds))
		}
		l := d.getLocation(s)
		if l < lowestLocation {
			lowestLocation = l
		}
	}
	return int(lowestLocation)
}

func (d *Day5) getLocation(s int) int {
	phase := "seed"
	value := s
	for {
		m := d.maps[phase]
		value = m.get(value)
		phase = m.dest
		if phase == "location" {
			break
		}
	}
	return value
}

func (d *Day5) parse(lines []string) {
	d.maps = make(map[string]*seedMap, 0)
	curMap := (*seedMap)(nil)
	regex := regexp.MustCompile(".*-to-.* map:")

	for _, line := range lines {
		// First parse the seeds
		if strings.Contains(line, "seeds:") {
			seedStrs := strings.Fields(strings.TrimPrefix(line, "seeds: "))
			for _, str := range seedStrs {
				seed, _ := strconv.Atoi(str)
				d.seeds = append(d.seeds, seed)
			}
		}

		match := regex.Match([]byte(line))
		if match {
			if curMap != nil {
				d.maps[curMap.src] = curMap
			}
			srcDest := strings.Split(strings.TrimSuffix(line, " map:"), "-to-")
			curMap = &seedMap{src: srcDest[0], dest: srcDest[1]}
		} else {
			if curMap != nil && strings.TrimSpace(line) != "" {
				fields := strings.Fields(line)
				ds, _ := strconv.Atoi(fields[0])
				ss, _ := strconv.Atoi(fields[1])
				rl, _ := strconv.Atoi(fields[2])
				curMap.addRange(seedRange{destStart: ds, sourceStart: ss, rangeLen: rl})
			}
		}
	}
	d.maps[curMap.src] = curMap
}

func (d *Day5) getSeeds() []int {
	seeds := make([]int, 0)
	for i := 0; i < len(d.seeds)-1; i += 2 {
		for j := d.seeds[i]; j < d.seeds[i]+d.seeds[i+1]; j++ {
			seeds = append(seeds, j)
		}
	}
	return seeds
}

type seedRange struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

type seedMap struct {
	src, dest string
	ranges    []seedRange
}

func (sm *seedMap) addRange(sr seedRange) {
	sm.ranges = append(sm.ranges, sr)
}

func (sm *seedMap) get(srcVal int) int {
	for _, sr := range sm.ranges {
		if srcVal >= sr.sourceStart && srcVal <= sr.sourceStart+sr.rangeLen {
			return sr.destStart + (srcVal - sr.sourceStart)
		}
	}
	return srcVal
}

func init() {
	solvers[5] = &Day5{}
}
