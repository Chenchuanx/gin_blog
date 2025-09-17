// 修复api.js文件，确保所有函数都正确导出
// 首先确保文件开头添加了所有必要的导入
import { getUser } from './utils.js';

const API_BASE_URL = 'http://127.0.0.1:8080';

// 然后确保所有函数都有export关键字
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

// 新增文章相关API函数

/**
 * 获取当前用户的文章列表
 * @returns {Promise<Object>} - 文章列表数据
 */
export async function getMyArticles() {
  try {
    const user = getUser();
    if (!user) {
      throw { error: '用户未登录' };
    }
    
    const response = await fetch(`${API_BASE_URL}/articles/getmy`, {
      method: 'POST',  // 改为POST请求
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        author_id: user.id  // 发送当前登录用户的ID
      })
    });
    
    const responseData = await response.json();
    
    if (responseData.code !== "0") {
      throw { error: responseData.msg || '获取我的文章失败' };
    }
    
    return responseData;
  } catch (error) {
    console.error('获取我的文章失败:', error);
    throw error;
  }
}

/**
 * 获取所有文章列表
 * @returns {Promise<Object>} - 文章列表数据
 */
export async function getAllArticles() {
  try {
    // 修改：使用新的路由
    const response = await fetch(`${API_BASE_URL}/articles/get`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    
    const responseData = await response.json();
    
    if (responseData.code !== "0") {
      throw { error: responseData.msg || '获取文章列表失败' };
    }
    
    return responseData;
  } catch (error) {
    console.error('获取文章列表失败:', error);
    throw error;
  }
}

/**
 * 创建文章
 * @param {Object} article - 文章数据
 * @returns {Promise<Object>} - 创建结果
 */
export async function createArticle(article) {
  try {
    // 修改：使用新的路由
    const data = await sendRequest('/articles/create', article);
    return data;
  } catch (error) {
    throw error;
  }
}

/**
 * 更新文章
 * @param {Object} article - 文章数据
 * @returns {Promise<Object>} - 更新结果
 */
export async function updateArticle(article) {
  try {
    // 修改：使用带ID的新路由
    const response = await fetch(`${API_BASE_URL}/articles/update/${article.id}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(article)
    });
    
    const responseData = await response.json();
    
    if (responseData.code !== "0") {
      throw { error: responseData.msg || '更新文章失败' };
    }
    
    return responseData;
  } catch (error) {
    console.error('更新文章失败:', error);
    throw error;
  }
}

/**
 * 删除文章
 * @param {number} id - 文章ID
 * @returns {Promise<Object>} - 删除结果
 */
export async function deleteArticle(id) {
  try {
    // 修改：使用带ID的新路由
    const response = await fetch(`${API_BASE_URL}/articles/delete/${id}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    
    const responseData = await response.json();
    
    if (responseData.code !== "0") {
      throw { error: responseData.msg || '删除文章失败' };
    }
    
    return responseData;
  } catch (error) {
    console.error('删除文章失败:', error);
    throw error;
  }
}

/**
 * 获取文章详情
 * @param {number} id - 文章ID
 * @returns {Promise<Object>} - 文章详情数据
 */
export async function getArticleDetail(id) {
  try {
    const response = await fetch(`${API_BASE_URL}/articles/get/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    
    const responseData = await response.json();
    
    if (responseData.code !== "0") {
      throw { error: responseData.msg || '获取文章详情失败' };
    }
    
    return responseData;
  } catch (error) {
    console.error('获取文章详情失败:', error);
    throw error;
  }
}

// 确认changeArticleLike函数的导出语法正确
/**
 * 更改文章点赞数
 * @param {number} id - 文章ID
 * @param {number} count - 点赞数变化值（1或-1）
 * @returns {Promise<Object>} - 操作结果
 */
export async function changeArticleLike(id, count) {
  try {
    const response = await fetch(`${API_BASE_URL}/articles/like/${id}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ count })
    });
    
    const responseData = await response.json();
    
    if (responseData.code !== "0") {
      throw { error: responseData.msg || '点赞操作失败' };
    }
    
    return responseData;
  } catch (error) {
    console.error('点赞操作失败:', error);
    throw error;
  }
}