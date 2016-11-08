package main

import "math"

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	theta := lon1 - lon2
	dist := math.Sin(deg2rad(lat1))*math.Sin(deg2rad(lat2)) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Cos(deg2rad(theta))
	dist = math.Acos(dist)
	dist = rad2deg(dist)
	dist = dist * 60 * 1.1515

	return dist
}

func deg2rad(deg float64) float64 {
	return (deg * math.Pi / 180.0)
}

func rad2deg(rad float64) float64 {
	return (rad / math.Pi * 180.0)
}
