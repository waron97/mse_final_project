import { Button, Card, Input, Typography } from "antd";
import classNames from "classnames";
import { FC, useState } from "react";
import { AiOutlineSearch } from "react-icons/ai";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";

import { Theme } from "@theme";

const { Title, Paragraph } = Typography;

// ----------------------------------------------------------------------------

interface LandingProps {
  className?: string;
}

const _Landing: FC<LandingProps> = (props) => {
  // -------------------------------------
  // Props destructuring
  // -------------------------------------

  const { className } = props;

  // -------------------------------------
  // Hooks (e.g. useState, useMemo ...)
  // -------------------------------------

  const [input, setInput] = useState<string>("");
  const navigate = useNavigate();

  // -------------------------------------
  // Effects
  // -------------------------------------

  // -------------------------------------
  // Component functions
  // -------------------------------------

  function handleSearch(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    navigate(`/search?query=${input}`);
  }

  function renderSearchForm() {
    return (
      <div>
        <Title className="title">MSE</Title>
        <Paragraph className="text">
          Aron Winkler, Nino Meisinger, Qin Gu, Zhiyuan Liu
        </Paragraph>
        <form onSubmit={handleSearch} className="search-form">
          <Input
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Input your query"
          ></Input>
          <Button type="primary" shape="circle" htmlType="submit">
            <AiOutlineSearch />
          </Button>
        </form>
      </div>
    );
  }

  // -------------------------------------
  // Component local variables
  // -------------------------------------

  return (
    <div className={classNames([className])}>
      <div className="main">
        <Card>{renderSearchForm()}</Card>
      </div>
    </div>
  );
};

// ----------------------------------------------------------------------------

const Landing = styled(_Landing)<Theme>`
  & {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 50px 0px;
    width: 100vw;
    min-height: 100vh;

    .main {
      max-width: 600px;
      width: 100%;

      .title {
        text-align: center;
        margin-bottom: 6px;
      }

      .text {
        text-align: center;
        margin-bottom: 24px;
      }

      .search-form {
        display: flex;
        gap: 12px;
        align-items: center;

        svg {
          font-size: 20px;
        }
      }
    }
  }
`;

export default Landing;
