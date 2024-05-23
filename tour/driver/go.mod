module steebe.dev/tour/driver

go 1.22.0

replace steebe.dev/tour/pointers => ../pointers

require (
	steebe.dev/tour/functions v0.0.0-00010101000000-000000000000
	steebe.dev/tour/pointers v0.0.0-00010101000000-000000000000
	steebe.dev/tour/slices v0.0.0-00010101000000-000000000000
	steebe.dev/tour/structs v0.0.0-00010101000000-000000000000
)

replace steebe.dev/tour/structs => ../structs

replace steebe.dev/tour/slices => ../slices

replace steebe.dev/tour/functions => ../functions
