document.addEventListener("DOMContentLoaded", function () {
  const addOrderBtn = document.getElementById("addOrderBtn");
  const addCompleteOrderBtn = document.getElementById("addCompleteOrderBtn");
  const addClientBtn = document.getElementById("addClientBtn");
  const editClientBtn = document.getElementById("editClientBtn");
  const deleteClientBtn = document.getElementById("deleteClientBtn");
  const showOrdersBtn = document.getElementById("showOrdersBtn");
  const showClientsBtn = document.getElementById("showClientsBtn");

  const showClientsList = document.getElementById("showClientsList");
  const showOrdersList = document.getElementById("showOrdersList");
  const orderModal = document.getElementById("orderModal");
  const completeOrderModal = document.getElementById("completeOrderModal");
  const addClientModal = document.getElementById("addClientModal");
  const editClientModal = document.getElementById("editClientModal");
  const deleteClientModal = document.getElementById("deleteClientModal");

  function showModal(modal) {
    modal.style.display = "block";
  }

  function hideAllModals() {
    orderModal.style.display = "none";
    completeOrderModal.style.display = "none";
    addClientModal.style.display = "none";
    editClientModal.style.display = "none";
    deleteClientModal.style.display = "none";
    showOrdersList.style.display = "none";
  }

  showClientsBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(showClientsList);
  });

  showOrdersBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(showOrdersList);
  });

  addOrderBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(orderModal);
  });

  addCompleteOrderBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(completeOrderModal);
  });

  addClientBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(addClientModal);
  });

  editClientBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(editClientModal);
  });

  deleteClientBtn.addEventListener("click", function () {
    hideAllModals();
    showModal(deleteClientModal);
  });

  const closeButtons = document.querySelectorAll(".close");

  closeButtons.forEach(function (button) {
    button.addEventListener("click", function () {
      const modal = button.parentElement.parentElement;
      modal.style.display = "none";
    });
  });

  window.addEventListener("click", function (event) {
    if (event.target.classList.contains("modal")) {
      event.target.style.display = "none";
    }
  });
});

