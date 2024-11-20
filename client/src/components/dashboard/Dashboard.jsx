import useCurrentUser from "@hooks/useCurrentUser.js";
import DiscordLogin from "@components/DiscordLogin";

export default function Dashboard() {
	const { data: user, isLoading, isError } = useCurrentUser();

	if (isError) return <p>Error loading user session</p>;

	return (
		<div class="text-3xl">
			<p class="text-red-500 text-3xl">this is the dashboard page!</p>
			{isLoading ? <p>loading...</p> : (
				<h3>
					Welcome, {user.Name}!
				</h3>
			)}
      <br />
      <br />
      <br />
      {
        isLoading ? <p>loading...</p> : (
        !!user.DiscordId ? <p>you are connected to discord {user.DiscordId}</p> : <DiscordLogin />
        )
      }
		</div>
	);
}
