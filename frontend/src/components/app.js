import { h } from 'preact';
import { Router } from 'preact-router';

import Header from './header';

// Code-splitting is automated for `routes` directory
import Home from '../routes/home';
import Profile from '../routes/profile';

import Login from '../routes/login';
import Dashboard from '../routes/dashboard';

import Admin from '../routes/admin';
import Volunteer from '../routes/volunteer';

const App = () => (
	<div id="app">
		<Header />
		<main>
			<Router>
				<Home path="/" />
				<Profile path="/profile/" user="me" />
				<Profile path="/profile/:user" />

				{
					// Normal to everyone (including admins, hackers, and volunteers)
				}
				<Login path="/login" />
				<Dashboard path="/dashboard" />


				{
					// Special Permissions (not available to hackers)
				}
				<Admin path="/admin" />
				<Volunteer path="/volunteer" />
			</Router>
		</main>
	</div>
);

export default App;
