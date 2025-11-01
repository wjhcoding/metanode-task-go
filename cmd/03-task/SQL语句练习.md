# 03-task

## 题目：基本 CRUD 操作（students 表）
表结构说明：
- id: 主键，自增
- name: 学生姓名（字符串）
- age: 学生年龄（整数）
- grade: 学生年级（字符串）

### 要求与实现

1. 向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。

```sql
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');
```

2. 查询 students 表中所有年龄大于 18 岁的学生信息。

```sql
SELECT * FROM students WHERE age > 18;
```

3. 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。

```sql
UPDATE students SET grade = '四年级' WHERE name = '张三';
```

4. 删除 students 表中年龄小于 15 岁的学生记录。

```sql
DELETE FROM students WHERE age < 15;
```



## 题目：账户转账事务（MySQL）

表结构（示例）：
- accounts(id PRIMARY KEY AUTO_INCREMENT, balance DECIMAL(18,2))
- transactions(id PRIMARY KEY AUTO_INCREMENT, from_account_id INT, to_account_id INT, amount DECIMAL(18,2), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)

要求：实现从账户 A 向账户 B 转账 100 元的事务逻辑，若账户 A 余额不足则回滚。

### 实现（推荐：存储过程 + 行锁）
下面为 MySQL 存储过程实现，示例中转账金额使用参数传入，调用时传入 A、B 的 id 和金额（例如 1 -> 2, 100）。

```sql
DELIMITER //
CREATE PROCEDURE transfer_funds(IN p_from INT, IN p_to INT, IN p_amount DECIMAL(18,2))
BEGIN
  DECLARE v_balance DECIMAL(18,2);
  -- 开始事务
  START TRANSACTION;
  -- 锁定转出账户行，防止并发修改
  SELECT balance INTO v_balance FROM accounts WHERE id = p_from FOR UPDATE;
  IF v_balance IS NULL THEN
    ROLLBACK;
    SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'From account not found';
  ELSEIF v_balance < p_amount THEN
    ROLLBACK;
    SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Insufficient funds';
  ELSE
    UPDATE accounts SET balance = balance - p_amount WHERE id = p_from;
    UPDATE accounts SET balance = balance + p_amount WHERE id = p_to;
    INSERT INTO transactions (from_account_id, to_account_id, amount)
      VALUES (p_from, p_to, p_amount);
    COMMIT;
  END IF;
END //
DELIMITER ;
```

调用示例：
```sql
CALL transfer_funds(1, 2, 100);
```

### 其他实现（单次事务 SQL 示例）
如果不想创建存储过程，也可在客户端（应用）中执行以下 SQL（确保在同一事务内执行）：

```sql
START TRANSACTION;
SELECT balance FROM accounts WHERE id = 1 FOR UPDATE;
-- 检查返回的 balance 是否 >= 100（由应用逻辑判断）
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;
INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (1, 2, 100);
COMMIT;
-- 若余额不足则执行 ROLLBACK;
```

### 注意事项
- accounts.balance 建议使用 DECIMAL 类型，避免浮点误差。
- 使用 SELECT ... FOR UPDATE 锁定行以防并发冲突。
- 在高并发场景下可考虑对 accounts 表做更细粒度的保障与幂等处理（如事务重试、唯一事务记录等