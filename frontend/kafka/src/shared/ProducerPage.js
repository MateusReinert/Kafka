import React, { useState } from 'react';
import './ProducerPage.css'; // Importe o CSS

function ProducerPage() {
  // Estado para armazenar o valor da mensagem
  const [inputMessage, setInputMessage] = useState('');
  const [alertMessage, setAlertMessage] = useState('');

  // Função que envia a mensagem via POST
  const sendMessage = async () => {
    try {
      // Realizando a requisição POST com a mensagem
      const response = await fetch(process.env.REACT_APP_REST_HOST+'/send-message', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: inputMessage }), // Enviando a mensagem no corpo
      });

      // Se a resposta for bem-sucedida, limpa o input e mostra um alerta
      if (response.ok) {
        setInputMessage(''); // Limpar o campo de entrada

        // Formatando a mensagem para exibição bonita
        const formattedMessage = `Producer:\nMensagem:\n${inputMessage}`;

        setAlertMessage(formattedMessage); // Exibe a mensagem formatada
      } else {
        // Caso contrário, exibe um alerta de erro
        setAlertMessage('Erro ao enviar mensagem');
      }
    } catch (error) {
      // Exibe um erro caso a requisição falhe
      setAlertMessage('Erro ao enviar mensagem');
    }
  };

  return (
    <div className="container">
      <h1>Criar Mensagem</h1>
      
      {/* Input para o usuário escrever a mensagem */}
      <input
        type="text"
        value={inputMessage}
        onChange={(e) => setInputMessage(e.target.value)} // Atualiza o estado com o valor do input
        placeholder="Escreva sua mensagem"
      />

      {/* Botão para enviar a mensagem */}
      <button onClick={sendMessage}>Enviar Mensagem</button>

      {/* Exibir o alerta de sucesso ou erro */}
      {alertMessage && (
        <div className={alertMessage.includes('Erro') ? 'alert error' : 'alert'}>
          {alertMessage}
        </div>
      )}
    </div>
  );
}

export default ProducerPage;
