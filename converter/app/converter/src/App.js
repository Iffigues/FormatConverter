import React, { useState, useCallback}  from 'react';
import MultipleFileUpload from './component/uplaod';
import CheckboxForm from './component/form'
import ServerHandler from './component/serverHandler'
import './App.css';

function App() {
	const [selectedOptions, setSelectedOptions] = useState([]);
	const [acceptedFiles, setAcceptedFiles] = useState([]);

	const handleCheckboxChange = (value) => {
		if (selectedOptions.includes(value)) {
			setSelectedOptions(selectedOptions.filter((option) => option !== value));
		} else {
			setSelectedOptions([...selectedOptions, value]);
		}
	};
	
	const handleDrop = useCallback((files) => {
		setAcceptedFiles(files);
	}, [setAcceptedFiles]);


	const handleSubmit = (e) => {
		e.preventDefault();
		console.log('Selected Options:', selectedOptions);
	};


  	return (
		<div>
			<MultipleFileUpload
				onDrop={handleDrop}
			/>
			<CheckboxForm
				selectedOptions={selectedOptions}
				handleCheckboxChange={handleCheckboxChange}
				handleSubmit={handleSubmit}
			/>
			<ServerHandler selectedOptions={selectedOptions} files={acceptedFiles} />
		</div>
	);
}

export default App;
