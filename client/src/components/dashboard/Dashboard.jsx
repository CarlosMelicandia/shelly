import useCurrentUser from "@hooks/useCurrentUser.js";

export default function Dashboard() {
	const { data: user, isLoading, isError } = useCurrentUser();

	if (isError) return <p>Error loading user session</p>;

	return (
		<div>
			<p class="text-red-500">this is the dashboard page!</p>
			{isLoading ? <p>loading...</p> : (
				<h1>
					Welcome, this is your token:{" "}
					{user
						? user.token
						: "you dont have a token"}
				</h1>
			)}
		</div>
	);
}
