const Login = () => {
	const handleLogin = () => {
		// Redirect the user to the Go backend's /login route to start Google OAuth
		window.location.href = "http://localhost:8000/api/auth/login";
	};

	return (
		<button onClick={handleLogin} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
			Sign in with Google
		</button>
	);
};

export default Login;
