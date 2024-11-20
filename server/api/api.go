package api

type User struct {
  Id string
  UserId string
  Name string
  FamilyName *string
  Email string
  DiscordId *string
  HackerId *int
  IsAdmin bool
  IsVolunteer bool
  IsMentor bool
  IsSponsor bool
}
