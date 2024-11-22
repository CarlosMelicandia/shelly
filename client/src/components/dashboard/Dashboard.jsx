import useCurrentHacker from "@hooks/useCurrentHacker.js";
import useCurrentUser from "@hooks/useCurrentUser.js";
import DiscordLogin from "@components/login/DiscordLogin";

export default function Dashboard() {
   const { data: hacker, isLoadingHacker } = useCurrentHacker();
   const { data: user, isLoadingUser } = useCurrentUser();

  if (isLoadingUser || isLoadingHacker) return <p>loading...</p>;
	if (!user || !hacker) return <p>You do not have a hacker application or aren't logged in!</p>;

	return (
		<div class="text-3xl">
			<p class="text-red-500 text-3xl">this is the dashboard page!</p>
			{isLoadingHacker || isLoadingUser ? <p>loading...</p> : (
				<h3>
					Welcome, {hacker?.first_name}!
				</h3>
			)}
      <br />
      <br />
      <br />
      {
        isLoadingUser ? <p>loading...</p> : (
        !!user?.discordId ? <p>you are connected to discord {user?.discordId}. Feel free to connect with another account <DiscordLogin /></p> : <DiscordLogin />
        )
      }
		</div>
	);
}
