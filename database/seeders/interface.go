package seeders

type Seeder interface  {
	Up() error
	Down() error
}
