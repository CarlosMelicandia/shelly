import { useState } from "preact/hooks"

/**
 * A reusable text input field component for various fields with validation.
 * Automatically renders a text input if "Other" is selected in a dropdown.
 *
 * @param {Object} props - The component props.
 * @param {Object} props.form - The React Hook Form instance containing `register` and `formState`.
 * @param {string} props.label - The label for the input field. Example: "Pronouns".
 * @param {string} props.name - The name and id for the input field, used for form handling. Example: "pronouns".
 * @param {string} [props.type="text"] - The input type. Example: "text", "email", "select".
 * @param {boolean} [props.isRequired=false] - Whether the input field is required. Defaults to `false`.
 * @param {Object} [props.validationRules={}] - Additional validation rules for the input.
 * @param {Array<string>} [props.options=[]] - Options for select inputs (if applicable).
 * @returns {JSX.Element} The rendered form input component.
 */
export default function TextField({
  form,
  label,
  name,
  type = "text",
  isRequired = false,
  validationRules = {},
  options = [],
}) {
  const {
    register,
    formState: { errors },
    setValue,
  } = form;

  const [isOtherSelected, setIsOtherSelected] = useState(false);

  const handleSelectChange = (e) => {
    const value = e.target.value;
    setIsOtherSelected(value === "Other");
    if (value !== "Other") {
      setValue(`${name}Other`, ""); // Clear the "Other" input if not needed
    }
  };

  return (
    <div className="mb-4">
      <label className="block text-gray-700 font-bold mb-2" htmlFor={name}>
        {isRequired ? `*${label}` : label}
      </label>
      {type === "select" ? (
        <>
          <select
            id={name}
            className={`w-full px-3 py-2 border rounded ${
              errors[name] ? "border-red-500" : "border-gray-300"
            }`}
            {...register(name, {
              required: isRequired ? `${label} is required` : false,
              ...validationRules,
            })}
            onChange={handleSelectChange}
          >
            <option value="">Select {label}</option>
            {options.map((option) => (
              <option key={option} value={option}>
                {option}
              </option>
            ))}
          </select>
          {isOtherSelected && (
            <div className="mt-2">
              <label
                className="block text-gray-700 font-bold mb-2"
                htmlFor={`other_${name}`}
              >
                Specify Other
              </label>
              <input
                id={`other_${name}`}
                type="text"
                className={`w-full px-3 py-2 border rounded ${
                  errors[`other_${name}`] ? "border-red-500" : "border-gray-300"
                }`}
                {...register(`other_${name}`, {
                  required: isOtherSelected ? "Please specify the other." : false,
                })}
              />
              {errors[`other_${name}`] && (
                <p className="text-red-500 text-sm mt-1">
                  {errors[`other_${name}`].message}
                </p>
              )}
            </div>
          )}
        </>
      ) : (
        <input
          id={name}
          type={type}
          className={`w-full px-3 py-2 border rounded ${
            errors[name] ? "border-red-500" : "border-gray-300"
          }`}
          {...register(name, {
            required: isRequired ? `${label} is required` : false,
            ...validationRules,
          })}
        />
      )}
      {errors[name] && <p className="text-red-500 text-sm mt-1">{errors[name].message}</p>}
    </div>
  );
}

