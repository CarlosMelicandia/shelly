import useCurrentUser from "@hooks/useCurrentUser.js";

export default function Dashboard() {
	const { data: user, isLoading, isError } = useCurrentUser();

	if (isError) return <p>Error loading user session</p>;

	return (
		<div class="text-3xl">
			<p class="text-red-500 text-3xl">this is the dashboard page!</p>
			{isLoading ? <p>loading...</p> : (
				<h3>
					Welcome, this is your token:{" "}
					{user
						? user.token
						: "you dont have a token"}
					<br />
					<br />
					And this is your user id: {" "}
					{user
						? user.userId
						: "you dont have a user id"}
				</h3>
			)}
		</div>
	);
}
