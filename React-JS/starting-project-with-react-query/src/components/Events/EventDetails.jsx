import { Link, Outlet } from "react-router-dom";

import Header from "../Header.jsx";
import { useMutation, useQuery } from "@tanstack/react-query";
import { deleteEvent, fetchEvent } from "../../utils/http.js";
import { queryClient } from "../../utils/http.js";
import { useState } from "react";
import Modal from "../UI/Modal.jsx";
import { useNavigate } from "react-router-dom";
import ErrorBlock from "../UI/ErrorBlock.jsx";
import { useParams } from "react-router-dom";

export default function EventDetails() {
  const params = useParams();
  const navigate = useNavigate();
  const [choose, setChoose] = useState(false);

  const { data, isPending, isError, error } = useQuery({
    queryKey: ["event", params.id],
    queryFn: ({ signal }) => fetchEvent({ id: params.id, signal }),
    gcTime: 30000,
  });

  const {
    mutate,
    isPending: isPendingDelete,
    isError: isErrorDelete,
    error: errorDelete,
  } = useMutation({
    mutationFn: () => deleteEvent({ id: params.id }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["events"],
        refetchType: "none",
      });
      navigate("/events");
    },
  });

  let content;

  if (isError) {
    content = (
      <div id="event-details-content" className="center">
        <ErrorBlock
          title="Failed to load event"
          message={
            error.info?.message ||
            "Failed to fetch events data, please try again later."
          }
        />
      </div>
    );
  }

  if (isPending) {
    content = (
      <div id="event-details-content" className="center">
        <p>Fetching event data...</p>
      </div>
    );
  }
  if (data) {
    const formattedDate = new Date(data.date).toLocaleDateString("en-US", {
      day: "numeric",
      month: "short",
      year: "numeric",
    });

    content = (
      <>
        <header>
          <h1>{data.title}</h1>
          <nav>
            <button onClick={() => setChoose((prev) => !prev)}>Delete</button>
            <Link to="edit">Edit</Link>
          </nav>
        </header>
        <div id="event-details-content">
          <img src={`http://localhost:3000/${data.image}`} />
          <div id="event-details-info">
            <div>
              <p id="event-details-location">{data.location}</p>
              <time dateTime={`Todo-DateT$Todo-Time`}>
                {formattedDate} @ {data.time}
              </time>
            </div>
            <p id="event-details-description">{data.description}</p>
          </div>
        </div>
      </>
    );
  }

  function onClose(event) {
    event.preventDefault();
    setChoose((prev) => !prev);
    navigate(`../events/${params.id}`);
  }

  return (
    <>
      {choose && (
        <Modal onClose={onClose}>
          {isPendingDelete ? (
            <p>Deleteing...</p>
          ) : (
            <>
              <p>
                Are You Sure Want to Delete Event <strong>{data.title}?</strong>
              </p>
              <p className="form-actions">
                <button
                  type="submit"
                  onClick={(event) => onClose(event)}
                  className="button-text"
                >
                  Cancel
                </button>
                <button type="submit" className="button" onClick={mutate}>
                  Delete
                </button>
              </p>
            </>
          )}
          {isErrorDelete && (
            <ErrorBlock
              title={"Failed to delete event"}
              message={
                errorDelete.info?.message ||
                "Failed to delete event, please try again later."
              }
            />
          )}
        </Modal>
      )}

      <Outlet />
      <Header>
        <Link to="/events" className="nav-item">
          View all Events
        </Link>
      </Header>
      <article id="event-details">{content}</article>
    </>
  );
}
