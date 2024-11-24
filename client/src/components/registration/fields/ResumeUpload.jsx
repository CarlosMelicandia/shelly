/**
 * A file upload component for uploading resumes.
 *
 * @param {Object} form - The React Hook Form instance containing `register` and `formState`.
 * @param {Function} form.register - Function to register the input for validation.
 * @param {Object} form.formState - The state of the form, including validation errors.
 * @param {Object} form.formState.errors - The errors object containing validation messages.
 * @returns {JSX.Element} The rendered resume upload component.
 */
export default function ResumeUpload({ form }) {
  const {
    register,
    formState: { errors },
  } = form;

  const _name = "resume";
  const _label = "Upload Resume";

  return (
    <div className="mb-4">
      <label className="block text-gray-700 font-bold mb-2" htmlFor={_name}>
        *{_label}
      </label>
      <input
        id={_name}
        type="file"
        accept=".pdf,.docx"
        className={`w-full px-3 py-2 border rounded ${errors[_name] ? "border-red-500" : "border-gray-300"
          }`}
        {...register(_name, {
          required: `${_label} is required`,
          validate: {
            isFileType: (fileList) => {
              const file = fileList[0];
              if (!file) return true; // No file uploaded, required handled separately
              const validExtensions = [".pdf", ".docx"];
              const fileExtension = file.name.slice(file.name.lastIndexOf(".")).toLowerCase();
              return (
                validExtensions.includes(fileExtension) ||
                "Resume must be a PDF or DOCX file"
              );
            },
            isFileSize: (fileList) => {
              const file = fileList[0];
              if (!file) return true; // No file uploaded, required handled separately
              const maxFileSize = 2 * 1024 * 1024; // 2MB
              return (
                file.size <= maxFileSize ||
                "File size must be less than 2MB"
              );
            },
          },
        })}
      />
      {errors[_name] && <p className="text-red-500 text-sm mt-1">{errors[_name].message}</p>}
    </div>
  );
}
