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
        json
      </label>
      <label>
        <input
          type="checkbox"
          value="yaml"
          checked={selectedOptions.includes("yaml")}
          onChange={() => handleCheckboxChange("yaml")}
        />
        yaml
      </label>
    </form>
  );
};

export default CheckboxForm
