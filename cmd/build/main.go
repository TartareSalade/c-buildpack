package main

import (
	"github.com/paketo-buildpacks/packit/v2"

)

func main(){

	packit.Detect(detect)

}

func detect(context packit.DetectContext) (packit.DetectResult, error){

	// TODO
	return packit.DetectResult{}, nil
}





