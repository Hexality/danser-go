package objects

import (
	"strconv"
	"math"
	"github.com/wieku/danser/bmath"
	"github.com/wieku/danser/audio"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/wieku/danser/render"
	"github.com/wieku/danser/settings"
)

const rpms = 0.00795

type Spinner struct {
	objData *basicData
	Timings *Timings
	sample  int
	rad     float64
	pos bmath.Vector2d
}

func NewSpinner(data []string) *Spinner {
	spinner := &Spinner{}
	spinner.objData = commonParse(data)
	spinner.objData.parseExtras(data, 6)

	spinner.objData.EndTime, _ = strconv.ParseInt(data[5], 10, 64)

	sample, _ := strconv.ParseInt(data[4], 10, 64)

	spinner.sample = int(sample)

	spinner.objData.EndPos = spinner.objData.StartPos
	return spinner
}

func (self *Spinner) GetBasicData() *basicData {
	return self.objData
}

func (self *Spinner) GetPosition() bmath.Vector2d {
	return self.pos
}

func (self *Spinner) SetTiming(timings *Timings) {
	self.Timings = timings
}

func (self *Spinner) Update(time int64) bool {
	if time < self.objData.EndTime {
		self.rad = rpms * float64(time-self.objData.StartTime) * 2 * math.Pi
		self.pos = bmath.NewVec2dRad(self.rad, 10).Add(self.objData.StartPos)
		return false
	}

	index := self.objData.customIndex

	if index == 0 {
		index = self.Timings.Current.SampleIndex
	}

	if self.objData.sampleSet == 0 {
		audio.PlaySample(self.Timings.Current.SampleSet, self.objData.additionSet, self.sample, index, self.Timings.Current.SampleVolume)
	} else {
		audio.PlaySample(self.objData.sampleSet, self.objData.additionSet, self.sample, index, self.Timings.Current.SampleVolume)
	}

	return true
}

func (self *Spinner) Render(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch) bool {
	alpha := 1.0

	if time < self.objData.StartTime-int64(preempt)/2 {
		alpha = float64(time-(self.objData.StartTime-int64(preempt))) / (preempt / 2)
	} else if time >= self.objData.EndTime {
		alpha = 1.0 - float64(time-self.objData.EndTime)/(preempt/2)
	} else {
		alpha = float64(color[3])
	}

	batch.SetTranslation(self.objData.StartPos)

	if settings.DIVIDES >= settings.Objects.MandalaTexturesTrigger {
		alpha *= settings.Objects.MandalaTexturesAlpha
	}

	batch.SetColor(1, 1, 1, alpha)//float64(color[0]), float64(color[1]), float64(color[2]), alpha)
	batch.SetScale(1, 1)

	batch.SetRotation(self.rad)
	batch.SetSubScale(20,20)


	batch.DrawUnitR(3)
	batch.DrawUnitR(4)

	scl := 16 + math.Min(220, math.Max(0, (1.0 - float64(time - self.objData.StartTime)/float64(self.objData.EndTime - self.objData.StartTime)) * 220))

	batch.SetSubScale(scl, scl)

	batch.DrawUnitR(5)

	batch.SetRotation(0)

	if time >= self.objData.EndTime+int64(preempt/2) {
		return true
	}

	return false
}