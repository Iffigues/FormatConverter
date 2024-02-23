import React  from 'react';

const CheckboxForm = ({ selectedOptions, handleCheckboxChange, handleSubmit }) => {
  return (
    <form onSubmit={handleSubmit}>
      <label>
        <input
          type="checkbox"
          value="option1"
          checked={selectedOptions.includes("option1")}
          onChange={() => handleCheckboxChange("option1")}
        />
        Option 1
      </label>
      <label>
        <input
          type="checkbox"
          value="option2"
          checked={selectedOptions.includes("option2")}
          onChange={() => handleCheckboxChange("option2")}
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
