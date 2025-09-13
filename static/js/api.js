// API基础URL
const API_BASE_URL = 'http://127.0.0.1:8080';

/**
 * 发送API请求的通用函数
 * @param {string} endpoint - API端点
 * @param {Object} data - 请求数据
 * @returns {Promise<Object>} - 返回响应数据
 */
async function sendRequest(endpoint, data) {
  try {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    
    // 总是尝试解析响应，无论状态码如何
    const responseData = await response.json();
    console.log('响应数据:', responseData);
    
    // 修改：同时检查字符串和数字类型的错误码
    if (responseData.code !== "0") {
      throw {
        error: responseData.msg || '请求失败',
        code: responseData.code,
        response: responseData
      };
    }
    
    return responseData;
  } catch (error) {
    console.error('API request failed:', error);
    throw error;
  }
}

/**
 * 用户登录
 * @param {string} username - 用户名
 * @param {string} password - 密码
 * @returns {Promise<Object>} - 登录结果
 */
export async function login(username, password) {
  try {
    const data = await sendRequest('/login', {
      username,
      password
    });
    // 验证响应数据完整性
    // if (!data || !data.user) {
    //   throw { error: '登录响应数据不完整' };
    // }
    return data;
  } catch (error) {
    throw error;
  }
}

/**
 * 用户注册
 * @param {string} username - 用户名
 * @param {string} password - 密码
 * @param {string} email - 邮箱
 * @returns {Promise<Object>} - 注册结果
 */
export async function register(username, password, email) {
  try {
    const data = await sendRequest('/sign_up', {
      username,
      password,
      email
    });
    return data;
  } catch (error) {
    // 直接传递原始错误，保留服务器返回的具体错误信息
    throw error;
  }
}

export async function changePassword(username, password, newPassword) {
  try {
    const data = await sendRequest('/change_password', {
      username,
      password,
      newpassword: newPassword
    });
    return data;
  } catch (error) {
    throw error;
  }
}