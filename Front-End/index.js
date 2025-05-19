async function carregarProdutos() {
  try {
    const resposta = await fetch("http://localhost:8000/products");
    const produtos = await resposta.json();

    const tabela = document.getElementById("lista-produtos");
    tabela.innerHTML = "";

    produtos.forEach(produto => {
      const linha = document.createElement("tr");

      linha.innerHTML = `
        <td>${produto.id}</td>
        <td>${produto.name}</td>
        <td>R$ ${parseFloat(produto.price).toFixed(2)}</td>
        <td>${produto.estoque}</td>
        <td>${produto.categoria}</td>
        <td>${produto.descricao}</td>
        <td>${produto.ativo ? "Sim" : "Não"}</td>
      `;

      tabela.appendChild(linha);
    });
  } catch (erro) {
    console.error("Erro ao carregar produtos:", erro);
  }
}

window.onload = carregarProdutos;
