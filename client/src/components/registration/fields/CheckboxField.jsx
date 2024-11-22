/**
 * A reusable checkbox input component for forms.
 *
 * @param {Object} props - The component props.
 * @param {Object} props.form - The React Hook Form instance containing `register` and `formState`.
 * @param {string} props.label - The label for the checkbox field. Example: "Agree to MLH News".
 * @param {string} props.name - The name for the checkbox field, used for form handling. Example: "agreedMLHNews".
 * @param {boolean} [props.isRequired=false] - Whether the checkbox is required. Defaults to `false`.
 * @returns {JSX.Element} The rendered checkbox input component.
 */
function CheckboxField({ form, label, name, isRequired = false }) {
  const {
    register,
    formState: { errors },
  } = form;

  return (
    <div className="mb-4">
      <label className="inline-flex items-center">
        <input
          type="checkbox"
          className={`form-checkbox h-5 w-5 text-blue-600 ${
            errors[name] ? "border-red-500" : "border-gray-300"
          }`}
          {...register(name, {
            required: isRequired ? `${label} is required` : false,
          })}
        />
        <span className="ml-2 text-gray-700">{label}</span>
      </label>
      {errors[name] && <p className="text-red-500 text-sm mt-1">{errors[name].message}</p>}
    </div>
  );
}

export default CheckboxField;

