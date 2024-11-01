module szechuansage/main

go 1.22.1

replace szechuansage/server => ../server/

replace szechuansage/slotgame => ../slotgame/

require (
	// github.com/mroth/weightedrand/v2 v2.1.0
	szechuansage/bonsai v0.0.0-00010101000000-000000000000
	szechuansage/server v0.0.0-00010101000000-000000000000
	szechuansage/slotgame v0.0.0-00010101000000-000000000000
)

replace szechuansage/bonsai => ../bonsai
