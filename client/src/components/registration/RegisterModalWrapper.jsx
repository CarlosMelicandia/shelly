import RegisterModal from "@components/registration/RegisterModal";
import useCurrentUser from "@hooks/useCurrentUser";
import useCurrentHacker from "@hooks/useCurrentHacker";

/**
 * We want to make sure the user has logged in with google and does not have an existing hacker application.
 * If they are not logged in or already has a hacker application tied to their userId, we want to not allow them to see the hacker application.
 *
 * @returns {JSX.Element} Conditional RegisterModal rendering
 */
const RegisterModalWrapper = () => {
  const { data: user, isLoadingUser, isErrorUser } = useCurrentUser();
  const { data: hacker, isLoadingHacker, isErrorHacker } = useCurrentHacker();

  if (isLoadingUser || isLoadingHacker) return

  const isUserLoggedIn = !isErrorUser && user;
  const hasHackerApp = !isErrorHacker && hacker;

  const cannotSeeForm = !!(!isUserLoggedIn || hasHackerApp)

  return cannotSeeForm ? null : <RegisterModal />
};

export default RegisterModalWrapper;

