import useCurrentUser from "@hooks/useCurrentUser.js";

const Admin = () => {
	const { data: user, isLoading, isError } = useCurrentUser();

  if (isLoading) return <p>loading...</p>
  if (isError) return <p>You are not logged in!</p>

  if (!user.is_admin) {
    return <p>you are not an admin</p>
  }

  return <p>secret admin stuff</p>

}

export default Admin
