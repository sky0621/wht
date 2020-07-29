package internal

import "time"

var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

var NilTime = time.Time{}
