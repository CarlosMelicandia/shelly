import { useEffect, useState } from "preact/hooks";
import RegisterForm from "./RegisterForm";

/*
 * We are wrapping the RegisterForm component with a Modal to make it easier to style the modal itself and have a separation of concerns -- meaning:
 * You either focus on the form itself or you focus on the container itself, ideally not both in one file
 * */
const RegisterModal = () => {
  const [isOpen, setIsOpen] = useState(false);

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    setIsOpen(params.get("register_form") === "true");
  }, []);

  const closeModal = () => setIsOpen(false);

  return (
    <>
      {isOpen && (
        <div className="fixed inset-0 flex justify-center items-center bg-gray-500 bg-opacity-50 z-50">
          <div className="bg-white p-6 rounded shadow-lg w-96">
            <button
              className="text-gray-500 float-right"
              onClick={closeModal}
            >
              &times;
            </button>
            <RegisterForm />
          </div>
        </div>
      )}
    </>
  );
};

export default RegisterModal;

