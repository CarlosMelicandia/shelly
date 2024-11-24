import { useForm } from "react-hook-form";
import TextField from "./fields/TextField";
import CheckboxField from "./fields/CheckboxField";
import ResumeUpload from "./fields/ResumeUpload";
import { useState } from "preact/hooks";

/**
 * React Hook Form for Hacker Applications.
 * This form is based on the `hacker_application` schema.
 *
 * @returns {JSX.Element} Hacker application form
 */
const HackerApplicationForm = () => {
  const form = useForm()
  const [isSubmitting, setIsSubmitting] = useState(false)
  const { handleSubmit } = form

  /**
   * Handles form submission.
   * @param {Object} data - Form data.
   */
  const onSubmit = async (data) => {
    if (isSubmitting) {
      return
    }
    setIsSubmitting(true)

    if (data.other_gender) {
      data.gender = data.other_gender
    }

    if (data.other_pronouns) {
      data.pronouns = data.other_pronouns
    }

    try {
      await fetch("http://localhost:8000/api/createHacker", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      window.location.replace("http://localhost:8000/dashboard")
    } catch (error) {
      console.error(error)
    }
    setIsSubmitting(false)
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-full max-w-lg bg-white shadow-md rounded p-6 overflow-y-auto max-h-screen"
        style={{ height: "90vh" }} // Ensure the form is scrollable
      >
        <h2 className="text-2xl font-bold text-center mb-4">Hacker Application</h2>

        <div className="space-y-4">
          <TextField
            form={form}
            label="First Name"
            name="first_name"
            type="text"
            isRequired
            validationRules={{
              minLength: { value: 2, message: "Enter your real name." },
              maxLength: { value: 100, message: "Enter your real name." },
            }}
          />
          <TextField
            form={form}
            label="Last Name"
            name="last_name"
            type="text"
            isRequired
            validationRules={{
              minLength: { value: 2, message: "Enter your real name." },
              maxLength: { value: 100, message: "Enter your real name." },
            }}
          />
          <TextField
            form={form}
            label="Age"
            name="age"
            type="number"
            isRequired
            validationRules={{
              min: { value: 18, message: "Age must be at least 18" },
              max: { value: 120, message: "Age cannot exceed 120" },
            }}
          />
          <TextField form={form} label="School" name="school" isRequired />
          <TextField form={form} label="Major" name="major" isRequired />
          <TextField
            form={form}
            label="Graduation Year"
            name="grad_year"
            type="number"
            isRequired
            validationRules={{
              min: { value: 1900, message: "Enter a valid year" },
              max: { value: new Date().getFullYear() + 10, message: "Year is too far in the future" },
            }}
          />
          <TextField
            form={form}
            label="Level of Study"
            name="level_of_study"
            type="text"
            isRequired
          />
          <TextField form={form} label="Country" name="country" isRequired />
          <TextField
            form={form}
            label="Email"
            name="email"
            type="email"
            isRequired
            validationRules={{
              pattern: {
                value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
                message: "Invalid email address",
              },
            }}
          />
          <TextField
            form={form}
            label="Phone Number"
            name="phone_number"
            type="tel"
            isRequired
            validationRules={{
              pattern: {
                value: /^\+?[0-9]{7,15}$/,
                message: "Invalid phone number",
              },
            }}
          />

          <TextField
            form={form}
            label="Test resume"
            name="resume_path"
            type="text"
            isRequired
          />

          <TextField
            form={form}
            label="Github"
            name="github"
            type="text"
          />

          <TextField
            form={form}
            label="Linkedin"
            name="linkedin"
            type="text"
          />

          <CheckboxField
            form={form}
            label="Are you an international student (currently on a non-immigrant visa status in the US such as F-1, or others)?"
            name="is_international"
          />

          <TextField
            form={form}
            label="Gender"
            name="gender"
            type="select"
            isRequired
            options={["Male", "Female", "Non-binary", "Other"]}
          />

          <TextField
            form={form}
            label="Pronouns"
            name="pronouns"
            isRequired
            options={["He/Him", "She/Her", "They/Them", "Other"]}
            type="select"
          />

          <TextField
            form={form}
            label="Ethnicity"
            name="ethnicity"
            isRequired
            options={["White", "Black", "Hispanic", "Other Options"]}
            type="select"
          />

          <CheckboxField
            form={form}
            label="Agree to MLH News"
            name="agreed_mlh_news"
          />
          <div className="flex justify-end">
            <button
              type="submit"
              className="bg-blue-500 text-white font-bold py-2 px-4 rounded hover:bg-blue-600"
              disabled={isSubmitting}
            >
              {isSubmitting ? <p>Submitting...</p> : <p>Submit</p>}
            </button>
          </div>
        </div>
      </form>
    </div>
  );
};

export default HackerApplicationForm;

