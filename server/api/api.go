package api

type User struct {
  Id string
  UserId string
  Name string
  FamilyName *string
  Email string
  DiscordUsername *string
  HackerId *int
  IsAdmin bool
  IsVolunteer bool
  IsMentor bool
  IsSponsor bool
}
