import QueryWrapper from "@wrappers/QueryWrapper"
import UserDashboard from "./UserDashboard"

export default function DashboardWrapper() {
	return (
		<QueryWrapper>
			<UserDashboard />
		</QueryWrapper>
	)
}
