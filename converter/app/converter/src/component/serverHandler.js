import React from 'react';

const ServerHandler = ({ selectedOptions, files }) => {
  const sendDataToServer = async () => {
    try {
console.log(files)
      const formData = new FormData();
      selectedOptions.forEach((option, index) => {
        formData.append(`options`, option);
      });

	files.forEach((file, index) => {
        formData.append(`files`, file);
      });


      // Make a POST request to the server
      const response = await fetch('http://gopiko.fr:8780/', {
        method: 'POST',
        body: formData,
      });

      // Handle the response as needed
      const result = await response.json();
      console.log('Server response:', result);
    } catch (error) {
      console.error('Error sending data to server:', error);
    }
  };

  return (
    <div>
      <button onClick={sendDataToServer}>Send Data to Server</button>
    </div>
  );
};

export default ServerHandler;
