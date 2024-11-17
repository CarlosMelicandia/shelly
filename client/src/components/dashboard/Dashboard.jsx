import useCurrentUser from "@hooks/useCurrentUser.js";

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
        !!user.DiscordUsername ? <p>you are connected to discord {user.DiscordUsername}</p> : <button>connect to discord (not working)</button>
        )
      }
		</div>
	);
}
