
function handleFileUpload(input, previewContainerId) {
  const previewContainer = document.getElementById(previewContainerId);

  Array.from(input.files).forEach(file => {
    const reader = new FileReader();

    reader.onload = function (e) {
      const preview = document.createElement('div');
      preview.className = 'flex items-center border rounded px-4 py-2 mb-2';
      preview.innerHTML = `
                <img src="${e.target.result}" alt="" class="w-12 h-12 mr-4">
                <p class="flex-1 text-gray-700">${file.name}</p>
                <button type="button" class="text-red-500" onclick="removePreview(this)">&times;</button>
            `;
      previewContainer.appendChild(preview);
    };

    reader.readAsDataURL(file);
  });
}

function removePreview(button) {
  button.parentElement.remove();
}

function enableDragAndDrop(dropAreaId, inputId, previewContainerId) {
  const dropArea = document.getElementById(dropAreaId);
  const input = document.getElementById(inputId);

  dropArea.addEventListener('dragover', (e) => {
    e.preventDefault();
    dropArea.classList.add('border-blue-500');
  });

  dropArea.addEventListener('dragleave', () => {
    dropArea.classList.remove('border-blue-500');
  });

  dropArea.addEventListener('drop', (e) => {
    e.preventDefault();
    dropArea.classList.remove('border-blue-500');

    const files = e.dataTransfer.files;
    const dataTransfer = new DataTransfer();

    // Append existing files
    Array.from(input.files).forEach(file => {
      dataTransfer.items.add(file);
    });

    // Append newly dropped files
    Array.from(files).forEach(file => {
      dataTransfer.items.add(file);
    });

    input.files = dataTransfer.files;
    handleFileUpload(input, previewContainerId);
  });

  input.addEventListener('change', () => {
    handleFileUpload(input, previewContainerId);
  });
}

// Enable drag-and-drop for banners and gallery
document.addEventListener('DOMContentLoaded', () => {
  enableDragAndDrop('banner-drop-area', 'banner-input', 'banner-preview');
  enableDragAndDrop('gallery-drop-area', 'gallery-input', 'gallery-preview');
});