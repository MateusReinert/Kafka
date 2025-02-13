import React, { useState, useEffect } from 'react';
import './MessagePage.css';  // Importe o CSS aqui

function MessagePage() {
  const [isConsuming, setIsConsuming] = useState(false);
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const socket = new WebSocket(process.env.REACT_APP_WS_HOST + '/ws');
  
    socket.onmessage = (event) => {
      let messageData;
      try {
        messageData = JSON.parse(event.data);
      } catch (error) {
        messageData = event.data;
      }
  
      setMessages((prevMessages) => [...prevMessages, messageData]);
    };
  
    socket.onerror = (error) => {
      console.error("Erro no WebSocket:", error);
    };
  
    return () => socket.close();
  }, []);

  const startConsumer = () => {
    fetch(process.env.REACT_APP_REST_HOST+'/start-consumer', { method: 'POST' })
      .then(() => setIsConsuming(true))
      .catch((error) => console.error('Erro ao iniciar o consumidor:', error));
  };

  const stopConsumer = () => {
    fetch(process.env.REACT_APP_REST_HOST+'/stop-consumer', { method: 'POST' })
      .then(() => setIsConsuming(false))
      .catch((error) => console.error('Erro ao parar o consumidor:', error));
  };

  return (
    <div className="container">
      <h1>Mensagens do Kafka</h1>
      <button onClick={isConsuming ? stopConsumer : startConsumer}>
        {isConsuming ? 'Parar Consumer' : 'Iniciar Consumer'}
      </button>
      <ul>
        {messages.length > 0 ? (
          messages.map((msg, index) => (
            <li key={index}>
              {/* Se a mensagem for um objeto, exibe o producer e a mensagem */}
              {typeof msg === 'object' ? (
                <>
                  <strong>Producer:</strong> {msg.producer || 'Desconhecido'} <br />
                  <strong>Mensagem:</strong> {msg.message || 'Sem conteúdo'}
                </>
              ) : (
                <>{msg}</>
              )}
            </li>
          ))
        ) : (
          <li className="no-messages">Nenhuma mensagem recebida ainda.</li>
        )}
      </ul>
    </div>
  );
}

export default MessagePage;
