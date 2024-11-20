const DiscordLogin = () => {
	const handleLogin = () => {
		window.location.href = "http://localhost:8000/api/auth/login/discord";
	};

	return (
		<button onClick={handleLogin} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
			Connect your discord account
		</button>
	);
};

export default DiscordLogin;

