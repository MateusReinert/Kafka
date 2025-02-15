import React, { useState } from 'react';
import './ProducerPage.css';

function ProducerPage() {
  const [inputMessage, setInputMessage] = useState('');
  const [alertMessage, setAlertMessage] = useState('');

  const sendMessage = async () => {
    try {
      const response = await fetch(process.env.REACT_APP_REST_HOST+'/send-message', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: inputMessage }),
      });

      if (response.ok) {
        setInputMessage('');

        const formattedMessage = `Producer:\nMensagem:\n${inputMessage}`;

        setAlertMessage(formattedMessage); 
      } else {
        setAlertMessage('Erro ao enviar mensagem');
      }
    } catch (error) {
      setAlertMessage('Erro ao enviar mensagem');
    }
  };

  return (
    <div className="container">
      <h1>Criar Mensagem</h1>
      
      <input
        type="text"
        value={inputMessage}
        onChange={(e) => setInputMessage(e.target.value)} 
        placeholder="Escreva sua mensagem"
      />

      <button onClick={sendMessage}>Enviar Mensagem</button>

      {alertMessage && (
        <div className={alertMessage.includes('Erro') ? 'alert error' : 'alert'}>
          {alertMessage}
        </div>
      )}
    </div>
  );
}

export default ProducerPage;
