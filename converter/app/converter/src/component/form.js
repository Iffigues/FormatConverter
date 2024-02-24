import React  from 'react';

const CheckboxForm = ({ selectedOptions, handleCheckboxChange, handleSubmit }) => {
  return (
    <form onSubmit={handleSubmit}>
      <label>
        <input
          type="checkbox"
          value="json"
          checked={selectedOptions.includes("json")}
          onChange={() => handleCheckboxChange("json")}
        />
        Option 1
      </label>
      <label>
        <input
          type="checkbox"
          value="yml"
          checked={selectedOptions.includes("yml")}
          onChange={() => handleCheckboxChange("yml")}
        />
        Option 2
      </label>
      <label>
        <input
          type="checkbox"
          value="option3"
          checked={selectedOptions.includes("option3")}
          onChange={() => handleCheckboxChange("option3")}
        />
        Option 3
      </label>
    </form>
  );
};

export default CheckboxForm
