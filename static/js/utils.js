/**
 * 显示错误消息
 * @param {string} message - 错误消息
 * @param {HTMLElement} element - 显示错误的元素
 */
export function showError(message, element) {
  element.textContent = message;
  element.style.display = 'block';
}

/**
 * 隐藏错误消息
 * @param {HTMLElement} element - 显示错误的元素
 */
export function hideError(element) {
  element.textContent = '';
  element.style.display = 'none';
}

/**
 * 显示成功消息
 * @param {string} message - 成功消息
 * @param {HTMLElement} element - 显示成功的元素
 */
export function showSuccess(message, element) {
  element.textContent = message;
  element.style.display = 'block';
}

/**
 * 隐藏成功消息
 * @param {HTMLElement} element - 显示成功的元素
 */
export function hideSuccess(element) {
  element.textContent = '';
  element.style.display = 'none';
}

/**
 * 保存用户信息到本地存储
 * @param {Object} user - 用户信息
 */
export function saveUser(user) {
  localStorage.setItem('user', JSON.stringify(user));
}

/**
 * 从本地存储获取用户信息
 * @returns {Object|null} - 用户信息或null
 */
export function getUser() {
  const userStr = localStorage.getItem('user');
  return userStr ? JSON.parse(userStr) : null;
}

/**
 * 从本地存储删除用户信息
 */
export function removeUser() {
  localStorage.removeItem('user');
}

/**
 * 检查用户是否已登录
 * @returns {boolean} - 是否已登录
 */
export function isAuthenticated() {
  return getUser() !== null;
}

/**
 * 验证邮箱格式
 * @param {string} email - 邮箱地址
 * @returns {boolean} - 是否有效
 */
export function isValidEmail(email) {
  const re = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return re.test(email);
}

/**
 * 禁用按钮并显示加载状态
 * @param {HTMLElement} button - 按钮元素
 * @param {string} loadingText - 加载状态文本
 */
export function disableButton(button, loadingText = '处理中...') {
  button.disabled = true;
  button.setAttribute('data-original-text', button.textContent);
  button.textContent = loadingText;
}

/**
 * 启用按钮并恢复原始文本
 * @param {HTMLElement} button - 按钮元素
 */
export function enableButton(button) {
  button.disabled = false;
  const originalText = button.getAttribute('data-original-text');
  if (originalText) {
    button.textContent = originalText;
    button.removeAttribute('data-original-text');
  }
}